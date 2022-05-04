package repositories

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sergiovenicio/graphql-students/graph/model"
)

type ChapterRepository struct {
	chapters []*model.Chapter
}

func (r *ChapterRepository) Create(name string) (*model.Chapter, error) {
	fmt.Println("[ChapterRepository] Create...")
	if name == "" {
		fmt.Println("[ChapterRepository] Error: name field is required")
		return nil, errors.New("name field is required")
	}

	id := uuid.NewV4()
	newChapter := &model.Chapter{
		ID:   id.String(),
		Name: name,
	}

	r.chapters = append(r.chapters, newChapter)
	fmt.Println("[ChapterRepository] Done")
	return newChapter, nil
}

func (r *ChapterRepository) GetChapters() ([]*model.Chapter, error) {
	fmt.Println("[ChapterRepository] GetChapters...")
	return r.chapters, nil
}
