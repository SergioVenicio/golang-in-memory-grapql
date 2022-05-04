package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/sergiovenicio/graphql-students/graph/model"
	"github.com/sergiovenicio/graphql-students/repositories"
)

type CategoryService struct {
	repository repositories.CategoryRepository
}

func (s *CategoryService) Create(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	fmt.Println("[CategoryService] Create...")
	return s.repository.Create(input.Name, *input.Description)
}

func (s *CategoryService) GetCategories(ctx context.Context) ([]*model.Category, error) {
	fmt.Println("[CategoryService] GetCategories...")
	return s.repository.GetCategories()
}

func (s *CategoryService) GetCategoryById(ctx context.Context, categoryId string) (*model.Category, error) {
	fmt.Println("[CategoryService] GetCategoryById...")
	categories, _ := s.repository.GetCategories()
	for _, category := range categories {
		if category.ID == categoryId {
			fmt.Println("[CategoryService] Done")
			return category, nil
		}
	}

	fmt.Println("[CategoryService] Error: category not found")
	return nil, errors.New("category not found")
}
