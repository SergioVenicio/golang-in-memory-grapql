package repositories

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sergiovenicio/graphql-students/graph/model"
)

type CourseRepository struct {
	courses []*model.Course
}

func (r *CourseRepository) Create(name string, desc string, categoryId string) (*model.Course, error) {
	fmt.Println("[CourseRepository] Create...")
	if name == "" {
		return nil, errors.New("name field is required")
	}

	id := uuid.NewV4()
	newCourse := &model.Course{
		ID:          id.String(),
		Name:        name,
		Description: &desc,
		CategoryID:  categoryId,
	}

	r.courses = append(r.courses, newCourse)
	fmt.Println("[CourseRepository] Done")
	return newCourse, nil
}

func (r *CourseRepository) GetCourses() ([]*model.Course, error) {
	fmt.Println("[CourseRepository] GetCourses...")
	return r.courses, nil
}
