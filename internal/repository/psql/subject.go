package psql

import (
	"GroupAssist/internal/domain"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	getSubjectAll     = `SELECT id, name FROM subjects ORDER BY id`
	getSubjectByID    = `SELECT id, name FROM subjects WHERE id=$1`
	createSubject     = `INSERT INTO subjects (name) VALUES ($1) RETURNING id, name`
	deleteSubjectByID = `DELETE FROM subjects WHERE id=$1`
	updateSubjectByID = `UPDATE subjects SET name=$1 WHERE id=$2`
)

type SubjectRepository struct {
	db *sqlx.DB
}

func NewSubjectRepository(db *sqlx.DB) *SubjectRepository {
	return &SubjectRepository{
		db: db,
	}
}

func (s *SubjectRepository) GetAll() ([]domain.Subject, error) {
	var subjects []domain.Subject
	err := s.db.Select(&subjects, getSubjectAll)
	return subjects, err
}

func (s *SubjectRepository) GetByID(id int) (domain.Subject, error) {
	var subject domain.Subject
	err := s.db.QueryRow(getSubjectByID, id).Scan(&subject.ID, &subject.SubjectName)
	return subject, err
}

func (s *SubjectRepository) Create(subject domain.Subject) (domain.Subject, error) {
	err := s.db.QueryRow(createSubject, subject.SubjectName).Scan(&subject.ID, &subject.SubjectName)
	log.Println(subject.SubjectName)
	return subject, err
}

func (s *SubjectRepository) Delete(id int) error {
	_, err := s.db.Exec(deleteSubjectByID, id)
	return err
}

func (s *SubjectRepository) Update(id int, subject domain.UpdateSubjectInput) error {
	_, err := s.db.Exec(updateSubjectByID, subject.SubjectName, id)
	return err
}
