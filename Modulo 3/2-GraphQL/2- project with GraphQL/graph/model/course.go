package model

// Course represents a course.
// Criei esse arquivo e adicionei essa struct que estava no models_gen.go
// OBS: No arquivo gqlgen.yml, adicionei o caminho para esse arquivo
// em course tenho que permanecer com o campo category
//type Course struct {
//	ID          string    `json:"id"`
//	Name        string    `json:"name"`
//	Description *string   `json:"description,omitempty"`
//	Category    *Category `json:"category"`
//}

// Retirei o campo Category *Category `json:"category"` pois ao rodar o comando generate de novo, ele faz o relacionamento sozinho
type Course struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
