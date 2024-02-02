package domain

type Subject struct {
	ID          int    `json:"id" db:"id"`
	SubjectName string `json:"name" db:"name" binding:"required"`
}

type UpdateSubjectInput struct {
	SubjectName string `json:"name" db:"name" binding:"required"`
}
