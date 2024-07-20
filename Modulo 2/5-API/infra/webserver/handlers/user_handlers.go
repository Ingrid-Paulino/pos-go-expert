package handlers

import (
	"net/http"
	"time"

	"github.com/Ingrid-Paulino/pos-go-expert/APIs/infra/database"
	"github.com/Ingrid-Paulino/pos-go-expert/APIs/internal/dto"
	"github.com/Ingrid-Paulino/pos-go-expert/APIs/internal/entity"
	"github.com/go-chi/jwtauth"
	"github.com/goccy/go-json"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
	//removi pois estou pegando do contexto da requisição o jwt
	//Jwt           *jwtauth.JWTAuth //vamos usar isso para gerar o token - esse já é pronto da lib go-chi
	JwtExperiesIn int
}

// func NewUserHandler(userDB database.UserInterface, Jwt *jwtauth.JWTAuth, JwtExperiesIn int) *UserHandler {
func NewUserHandler(userDB database.UserInterface, JwtExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB: userDB,
		//Jwt:           Jwt, //removi pois estou pegando do contexto da requisição o jwt
		JwtExperiesIn: JwtExperiesIn,
	}
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	// JWT: verifica se realmente é o usuario que está fazendo a requisição - se token autenticado é do usuario que está fazendo a requisição
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	//jwtExpiresIIn := r.Context().Value("jwt_expires_in").(int) //nn vai funcionar pois n estou passando no contexto e sim por injeção de dependencia
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(err)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
	//ou
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),                                                       //id do usuario
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExperiesIn)).Unix(), //expiração do token (é opcional)
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
