package user

import (
	"database/sql"

	"github.com/AJRDRGZ/db-query-builder/models"
	"github.com/AJRDRGZ/db-query-builder/postgres"
	sqlutil "github.com/alexyslozada/gosqlutils"

	"github.com/MikelSot/interseguro-challenge-auth/model"
)

const Table = "users"

var Fields = []string{
	"first_name",
	"last_name",
	"email",
	"password",
}

var (
	psqlInsert = postgres.BuildSQLInsert(Table, Fields)
	psqlGetAll = postgres.BuildSQLSelect(Table, Fields)
)

type User struct {
	db *sql.DB
}

func New(db *sql.DB) User {
	return User{db}
}

func (u User) Create(m model.User) error {
	stmt, err := u.db.Prepare(psqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.FirstName,
		m.Lastname,
		m.Email,
		m.Password,
	).Scan(&m.ID, &m.CreatedAt)

	return err
}

func (u User) GetWhere(specification models.FieldsSpecification) (model.User, error) {
	conditions, args := postgres.BuildSQLWhere(specification.Filters)
	query := psqlGetAll + " " + conditions

	query += " " + postgres.BuildSQLOrderBy(specification.Sorts)

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return model.User{}, err
	}
	defer stmt.Close()

	return u.scanRow(stmt.QueryRow(args...))
}

func (u User) scanRow(s sqlutil.RowScanner) (model.User, error) {
	m := model.User{}

	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.FirstName,
		&m.Lastname,
		&m.Email,
		&m.Password,
		&m.CreatedAt,
		&updatedAtNull,
	)

	m.UpdatedAt = updatedAtNull.Time

	return m, err
}
