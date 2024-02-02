package psql

import (
	"GroupAssist/internal/domain"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
)

var (
	getAllQueueBySubject = `
SELECT
	queue.id AS id,
	subjects.name AS subject_name,
	queue.title AS title,
	queue.subject_id AS subject_id,
	queue.is_open AS is_open,
	COUNT(queue_position.id) AS count
FROM
	queue
INNER JOIN
	subjects ON queue.subject_id = subjects.id
FULL OUTER JOIN
	queue_position ON queue.id = queue_position.queue_id
WHERE
	queue.subject_id = $1
GROUP BY
	queue.id, subjects.name
ORDER BY
	queue.id
	`

	getQueueByID = `
SELECT
	q.id,
	s.name as subject_name,
	q.title,
	q.subject_id,
	q.is_open,
	COUNT(qp.id) as count,
	COALESCE(json_agg(json_build_object(
	'entered_at', qp.entered_at, 'user', 'TODO')) FILTER(WHERE qp.id IS NOT NULL), '[]'::json) as position
FROM
	queue q
		INNER JOIN
	subjects s ON q.subject_id = s.id
		LEFT JOIN
	queue_position qp ON q.id = qp.queue_id
WHERE
	q.id = $1
GROUP BY
	q.id, s.name
	`
	createQueue     = `INSERT INTO queue (id, title, subject_id, is_open) VALUES ($1, $2, $3)`
	deleteQueueByID = `DELETE FROM queue WHERE id=$1`
	updateQueueByID = `UPDATE queue SET is_open=$1 WHERE id=$2`
)

type QueueRepository struct {
	db *sqlx.DB
}

func NewQueueRepository(db *sqlx.DB) *QueueRepository {
	return &QueueRepository{
		db: db,
	}
}

func (q *QueueRepository) GetAllBySubject(id int) ([]domain.QueuesResponse, error) {
	var queues []domain.QueuesResponse
	err := q.db.Select(&queues, getAllQueueBySubject, id)

	if len(queues) == 0 {
		return nil, errors.New("rows not found")
	}
	return queues, err
}

func (q *QueueRepository) GetByID(id int) (domain.QueueResponse, error) {
	var queue domain.QueueResponse
	var positionsRaw json.RawMessage

	subjectRow := q.db.QueryRow(getQueueByID, id)

	err := subjectRow.Scan(&queue.ID, &queue.SubjectName, &queue.Title,
		&queue.SubjectID, &queue.IsOpen, &queue.Count, &positionsRaw)
	if err != nil {
		return queue, err
	}

	err = json.Unmarshal(positionsRaw, &queue.Positions)
	if err != nil {
		return queue, err
	}
	return queue, err
}

func (q *QueueRepository) Create(queue domain.Queue) (domain.Queue, error) {
	err := q.db.QueryRow(createQueue, queue.Title, queue.SubjectID, queue.IsOpen).Scan(
		&queue.ID, &queue.Title, &queue.SubjectID, &queue.IsOpen)
	return queue, err
}

func (q *QueueRepository) Delete(id int) error {
	_, err := q.db.Exec(deleteQueueByID, id)
	return err
}

func (q *QueueRepository) Update(id int, queue domain.UpdateQueueInput) error {
	_, err := q.db.Exec(updateQueueByID, queue.IsOpen, id)
	return err
}
