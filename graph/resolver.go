package graph

import (
	"github.com/sergiovenicio/graphql-students/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryService services.CategoryService
	CourseService   services.CourseService
	ChapterService  services.ChapterService
}
