package model

type Course struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CategoryID  string
	ChapterIDs  []string
}

type NewCourse struct {
	Name        string  `json:"name"`
	CategoryID  string  `json:"categoryId"`
	Description *string `json:"description"`
}
