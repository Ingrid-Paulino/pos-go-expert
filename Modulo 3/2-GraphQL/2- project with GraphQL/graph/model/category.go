package model

// Category represents a category of courses.
// Criei esse arquivo e adicionei essa struct que estava no models_gen.go
// OBS: No arquivo gqlgen.yml, adicionei o caminho para esse arquivo
//type Category struct {
//	ID          string    `json:"id"`
//	Name        string    `json:"name"`
//	Description *string   `json:"description,omitempty"`
//	Courses     []*Course `json:"courses"`
//}

// Aqui tirei o  Courses     []*Course `json:"courses" pois ao rodar o comando generate de novo, ele faz o relacionamento sozinho
// generate será rodado depois que colocar os modelos separados no arquivo gqlgen.yml
type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
