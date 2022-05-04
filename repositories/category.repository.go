package repositories

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sergiovenicio/graphql-students/graph/model"
)

type CategoryRepository struct {
	categories []*model.Category
}

func (r *CategoryRepository) Create(name string, desc string) (*model.Category, error) {
	fmt.Println("[CategoryRepository] Create...")
	if name == "" {
		fmt.Println("[CategoryRepository] Erorr: name field is required")
		return nil, errors.New("name field is required")
	}

	id := uuid.NewV4()
	newCategory := &model.Category{
		ID:          id.String(),
		Name:        name,
		Description: &desc,
	}

	r.categories = append(r.categories, newCategory)
	fmt.Println("[CategoryRepository] Done")
	return newCategory, nil
}

func (r *CategoryRepository) GetCategories() ([]*model.Category, error) {
	fmt.Println("[CategoryRepository] GetCategories...")
	return r.categories, nil
}
