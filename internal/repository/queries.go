package repository

import (
	"database/sql"

	"github.com/DanilKlochkov/golang-graphql/internal/models"
	"github.com/DanilKlochkov/golang-graphql/pkg/db"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	database, err := db.Init()
	if err != nil {
		return nil, err
	}
	return &Repository{database}, nil
}

func (r *Repository) Get(id int64) (models.Note, error) {
	var m models.Note

	err := r.db.QueryRow(`SELECT * FROM note WHERE id = $1`, id).Scan(&m.ID, &m.Title, &m.Content)
	if err != nil {
		return models.Note{}, nil
	}

	return m, nil
}

func (r *Repository) Create(m models.Note) (models.Note, error) {
	var lastId int64
	query := `INSERT INTO note (title, content) VALUES ($1, $2) RETURNING id;`

	err := r.db.QueryRow(query, m.Title, m.Content).Scan(&lastId)
	if err != nil {
		return models.Note{}, err
	}

	model, err := r.Get(lastId)
	if err != nil {
		return models.Note{}, err
	}

	return model, nil
}
