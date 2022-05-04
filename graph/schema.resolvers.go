package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sergiovenicio/graphql-students/graph/generated"
	"github.com/sergiovenicio/graphql-students/graph/model"
)

func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	categoryId := obj.ID
	course := r.Resolver.CourseService.GetCoursesByCategoryId(ctx, categoryId)
	return course, nil
}

func (r *chapterResolver) Course(ctx context.Context, obj *model.Chapter) (*model.Course, error) {
	chapterId := obj.ID
	course := r.Resolver.CourseService.GetCourseByChapterId(ctx, chapterId)
	return course, nil
}

func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	fmt.Println("Category Resolver...")
	courseId := obj.ID
	course, _ := r.Resolver.CourseService.GetCourseById(ctx, courseId)
	category, _ := r.Resolver.CategoryService.GetCategoryById(ctx, course.CategoryID)
	return category, nil
}

func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	fmt.Println("Chapters Resolver...")
	courseId := obj.ID
	course, _ := r.Resolver.CourseService.GetCourseById(ctx, courseId)
	chapters, _ := r.Resolver.ChapterService.GetChaptersByIds(course.ChapterIDs)
	return chapters, nil
}

func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	return r.Resolver.CategoryService.Create(ctx, input)
}

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	return r.Resolver.CourseService.Create(ctx, input)
}

func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	fmt.Println("Create Chapter Resolver...")
	newChapter, _ := r.Resolver.ChapterService.Create(ctx, input)
	r.Resolver.CourseService.AddChapter(ctx, input.CourseID, newChapter.ID)
	return newChapter, nil
}

func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	return r.Resolver.CategoryService.GetCategories(ctx)
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.CourseService.GetCourses(ctx)
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.ChapterService.GetChapters(ctx)
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Chapter returns generated.ChapterResolver implementation.
func (r *Resolver) Chapter() generated.ChapterResolver { return &chapterResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type chapterResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
