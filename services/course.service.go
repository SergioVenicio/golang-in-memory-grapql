package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/sergiovenicio/graphql-students/graph/model"
	"github.com/sergiovenicio/graphql-students/repositories"
)

type CourseService struct {
	repository repositories.CourseRepository
}

func (s *CourseService) Create(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	fmt.Println("[CourseService] Create...")
	newCourse, err := s.repository.Create(input.Name, *input.Description, input.CategoryID)
	if err != nil {
		fmt.Printf("[CourseService] Error %v\n", err)
		return nil, err
	}

	fmt.Println("[CourseService] Done...")
	return newCourse, nil
}

func (s *CourseService) GetCourses(ctx context.Context) ([]*model.Course, error) {
	return s.repository.GetCourses()
}

func (s *CourseService) GetCourseById(ctx context.Context, courseId string) (*model.Course, error) {
	courses, _ := s.repository.GetCourses()
	for _, course := range courses {
		if course.ID == courseId {
			return course, nil
		}
	}

	return nil, errors.New("course not found")
}

func (s *CourseService) AddChapter(ctx context.Context, courseId string, chapterId string) (*model.Course, error) {
	course, _ := s.GetCourseById(ctx, courseId)
	course.ChapterIDs = append(course.ChapterIDs, chapterId)
	return nil, errors.New("course not found")
}

func (s *CourseService) GetCoursesByCategoryId(ctx context.Context, categoryId string) []*model.Course {
	var foundedCourses []*model.Course
	courses, _ := s.repository.GetCourses()
	for _, course := range courses {
		if course.CategoryID == categoryId {
			foundedCourses = append(foundedCourses, course)
		}
	}

	return foundedCourses
}

func (s *CourseService) GetCourseByChapterId(ctx context.Context, chapterId string) *model.Course {
	courses, _ := s.repository.GetCourses()
	for _, course := range courses {
		fmt.Println(course.ChapterIDs)
		for _, chapt := range course.ChapterIDs {
			if chapt == chapterId {
				return course
			}
		}
	}

	return nil
}
