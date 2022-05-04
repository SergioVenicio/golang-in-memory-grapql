package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/sergiovenicio/graphql-students/graph/model"
	"github.com/sergiovenicio/graphql-students/repositories"
)

type ChapterService struct {
	repository repositories.ChapterRepository
}

func (s *ChapterService) Create(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	fmt.Println("[ChapterService] Create...")
	newChapter, err := s.repository.Create(input.Name)
	if err != nil {
		fmt.Printf("[ChapterService] Error: %v", err)
		return nil, err
	}
	fmt.Println("[ChapterService] Done")
	return newChapter, nil
}

func (s *ChapterService) GetChapters(ctx context.Context) ([]*model.Chapter, error) {
	fmt.Println("[ChapterService] GetChapters...")
	return s.repository.GetChapters()
}

func (s *ChapterService) GetChapterById(chapterId string) (*model.Chapter, error) {
	fmt.Println("[ChapterService] GetChapterById...")
	chapters, _ := s.repository.GetChapters()
	for _, chapter := range chapters {
		if chapter.ID == chapterId {
			fmt.Println("[ChapterService] Done")
			return chapter, nil
		}
	}

	fmt.Println("[ChapterService] Error: chapter not found")
	return nil, errors.New("chapter not found")
}

func (s *ChapterService) GetChaptersByIds(chaptersIds []string) ([]*model.Chapter, error) {
	fmt.Println("[ChapterService] GetChaptersByIds...")
	var foundedChapters []*model.Chapter
	for _, cId := range chaptersIds {
		chapter, err := s.GetChapterById(cId)
		if err != nil {
			continue
		}

		foundedChapters = append(foundedChapters, chapter)
	}

	return foundedChapters, nil
}
