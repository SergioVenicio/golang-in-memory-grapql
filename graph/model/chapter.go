package model

type Chapter struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type NewChapter struct {
	Name     string `json:"name"`
	CourseID string `json:"courseId"`
}
