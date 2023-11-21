package author_db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/xlzpm/internal/author/model"
	"github.com/xlzpm/internal/author/storage"
	"github.com/xlzpm/pkg/client/postgresql"
	"github.com/xlzpm/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, author *model.Author) error {
	q := `INSERT INTO author 
				(name) 
		  VALUES 
		  		($1) 
		  RETURNING id`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}

		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context, sortOptions storage.SortOptions) ([]model.Author, error) {
	qb := sq.Select("id, name, age, is_alive, created_at").From("public.author")

	if sortOptions != nil {
		qb = qb.OrderBy(sortOptions.GetOrderBy())
	}

	sql, i, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(sql)))

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)

	for rows.Next() {
		var author model.Author

		err = rows.Scan(&author.ID, &author.Name, &author.Age, &author.IsAlive, &author.CreatedAT)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `SELECT id, name FROM author WHERE id = $1`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var ath model.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name)
	if err != nil {
		return model.Author{}, err
	}

	return ath, nil
}

func (r *repository) Update(ctx context.Context, author model.Author) error {
	q := `UPDATE author
		  	SET name = $1
		  WHERE id = $2
		  	RETURNING id, name`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	err := r.client.QueryRow(ctx, q, author.Name, author.ID).Scan(&author.ID, &author.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM author
		  WHERE id = $1
		  `

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q, id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
