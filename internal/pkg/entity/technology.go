package entity

import (
	"context"
	"database/sql"
	"time"
)

const timeoutInMillis int = 10

type technology struct {
	name    string
	version string
}

func (t technology) Name() string {
	return t.name
}

func (t technology) Version() string {
	return t.version
}

type repository interface {
	FindAll() ([]technology, error)
}

type repositoryImpl struct {
	db *sql.DB
}

func (r repositoryImpl) FindAll() ([]technology, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInMillis)*time.Millisecond)
	defer cancel()
	rows, err := r.db.QueryContext(ctx, "SELECT T.NAME, V.NUMBER FROM TECHNOLOGY T INNER JOIN VERSION V ON V.TECHNOLOGY_NAME = T.NAME")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var technologies []technology
	for rows.Next() {
		var tech technology
		err = rows.Scan(&tech.name, &tech.version)
		if err != nil {
			return nil, err
		}
		technologies = append(technologies, tech)
	}
	return technologies, nil
}

func NewRepository(db *sql.DB) repository {
	return repositoryImpl{db: db}
}
