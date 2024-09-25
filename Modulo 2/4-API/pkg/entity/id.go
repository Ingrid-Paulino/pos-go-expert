package entity

import "github.com/google/uuid"

//Esse arquivo pode ser compartilhado com outras pessoas que terão acesso ao projeto

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err

}
