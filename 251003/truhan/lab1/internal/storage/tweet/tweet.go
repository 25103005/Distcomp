package tweet

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Khmelov/Distcomp/251003/truhan/lab1/internal/model"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	ErrTweetNotFound     = fmt.Errorf("tweets not found")
	ErrInvalidTweetData  = fmt.Errorf("invalid tweet data")
	ErrInvalidForeignKey = fmt.Errorf("invalid foreign key passed")
)

type Repo interface {
	Create(ctx context.Context, req model.Tweet) (model.Tweet, error)
	GetList(ctx context.Context) ([]model.Tweet, error)
	Get(ctx context.Context, id int64) (model.Tweet, error)
	Update(ctx context.Context, req model.Tweet) (model.Tweet, error)
	Delete(ctx context.Context, id int64) error
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return repo{
		db: db,
	}
}

func (r repo) Create(ctx context.Context, req model.Tweet) (model.Tweet, error) {
	query := `INSERT INTO tweets (author_id, title, content) 
	          VALUES ($1, $2, $3) RETURNING id, created`

	var id int64
	var created time.Time

	err := r.db.QueryRowContext(ctx, query, req.AuthorID, req.Title, req.Content).Scan(&id, &created)
	if err != nil {
		log.Println(err)
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return model.Tweet{}, ErrTweetNotFound
		}

		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return model.Tweet{}, ErrInvalidTweetData
		}

		return model.Tweet{}, fmt.Errorf("failed to create tweet: %w", err)
	}

	req.ID = id
	req.Created = created

	return req, nil
}

func (r repo) GetList(ctx context.Context) ([]model.Tweet, error) {
	var result []model.Tweet

	query := `SELECT * FROM tweets`

	err := r.db.SelectContext(ctx, &result, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve result: %w", err)
	}

	if len(result) == 0 {
		return []model.Tweet{}, nil
	}

	return result, nil
}

func (r repo) Get(ctx context.Context, id int64) (model.Tweet, error) {
	var result model.Tweet

	query := `SELECT id, author_id, title, content, created, modified FROM tweets WHERE id = $1`

	err := r.db.GetContext(ctx, &result, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return result, ErrTweetNotFound
		}

		log.Println(err.Error())
		return result, fmt.Errorf("failed to retrieve tweet by ID: %w", err)
	}

	return result, nil
}

func (r repo) Update(ctx context.Context, req model.Tweet) (model.Tweet, error) {
	var result model.Tweet

	query := `UPDATE tweets SET author_id = $1, title = $2, content = $3, modified = $4 
	          WHERE id = $5 RETURNING id, author_id, title, content, created, modified`

	err := r.db.QueryRowContext(ctx, query, req.AuthorID, req.Title, req.Content, req.Modified, req.ID).
		Scan(&result.ID, &result.AuthorID, &result.Title, &result.Content, &result.Created, &result.Modified)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("tweet not found with id:", req.ID)
			return result, ErrTweetNotFound
		}

		log.Println("error with query:", err)
		return result, fmt.Errorf("failed to update tweet: %w", err)
	}

	return result, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tweets WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error executing DELETE query:", err)
		return fmt.Errorf("failed to delete tweet: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		log.Println("No tweet found with ID:", id)
		return ErrTweetNotFound
	}

	return nil
}
