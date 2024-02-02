package domain

type Queue struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	SubjectID int    `json:"subject_id"`
	IsOpen    bool   `json:"is_open"`
}

type UpdateQueueInput struct {
	IsOpen bool
}

type QueuesResponse struct {
	ID          int    `json:"id"`
	SubjectName string `json:"subject_name" db:"subject_name"`
	Title       string `json:"title"`
	SubjectID   int    `json:"subject_id" db:"subject_id"`
	IsOpen      bool   `json:"is_open" db:"is_open"`
	Count       int    `json:"count"`
}

type QueueResponse struct {
	ID          int             `json:"id"`
	SubjectName string          `json:"subject_name" db:"subject_name"`
	Title       string          `json:"title"`
	SubjectID   int             `json:"subject_id" db:"subject_id"`
	IsOpen      bool            `json:"is_open" db:"is_open"`
	Count       int             `json:"count"`
	Positions   []QueuePosition `json:"positions"`
}

type QueuePosition struct {
	User      string `json:"user"`
	EnteredAt string `json:"entered_at"`
}
