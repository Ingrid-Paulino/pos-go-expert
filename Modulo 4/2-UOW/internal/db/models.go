package db

type Category struct {
	ID   int32
	Name string
}

type Course struct {
	ID         int32
	Name       string
	CategoryID int32
}
