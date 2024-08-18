package service

import (
	"context"
	"io"

	"github.com/devfullcycle/gRPC/internal/database"
	"github.com/devfullcycle/gRPC/internal/pb"
)

type CategoryService struct {
	//imports necessários para implementar a interface pb.CategoryServiceServer
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	//implementação do método CreateCategory
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{ //obs: poderia retornar Category direto
		Category: categoryResponse,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	//implementação do método ListCategories
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesResponse []*pb.Category

	for _, category := range categories {
		categoryResponse := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}

		categoriesResponse = append(categoriesResponse, categoryResponse)
	}

	return &pb.CategoryList{Categories: categoriesResponse}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	//implementação do método ListCategories
	category, err := c.CategoryDB.Find(in.Id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

// Recebe todos os dados de uma vez e Retorna todos os dados de uma vez
func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}

	for {
		category, err := stream.Recv() // recebe os dados //vamos receber os dados que vamos pegar para criar a categoria
		if err == io.EOF {             //se chegou no final e  não tiver mais dados para receber, enviaremos todas as informaçoes dando um sendAndClose
			return stream.SendAndClose(&pb.CategoryList{Categories: categories.Categories})
		}
		if err != nil {
			return err
		}
		//Sera criado varias categorias e adicionado no slice de categorias
		//ao final, enviaremos todas as categorias criadas para o cliente
		//e fecharemos a conexão quando nn tiver mais dados para enviar
		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
	/*
		Passo a passo pra rodar a stream:
		1- rodar o servidor
		2- evans -r repl
		3- package pb
		4- service CategoryService
		5- call CreateCategoryStream
		6- digitar os dados da categoria
		7- para sair ctrl+D
		8- Sera listado todas as categorias criadas
	*/
}

// Recebo um dado(categoria) e retorno um dado(categoria), vou enviando e vou recebendo
func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv() //recebe os dados na stream
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		categoryResult, err := c.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		}) //envia os dados da stream

		if err != nil {
			return err
		}
	}
	/*
		Passo a passo pra rodar a stream:
		1- rodar o servidor
		2- evans -r repl
		3- package pb
		4- service CategoryService
		5- call CreateCategoryStreamBidirectional
		6- digita o dado da categoria
		7- Recebe a categoria criada/resultado
		8- digita o dado da categoria
		9- Recebe a categoria criada/resultado
		10- digita o dado da categoria
		11- Recebe a categoria/resultado
		12- para sair ctrl+D
	*/
}
