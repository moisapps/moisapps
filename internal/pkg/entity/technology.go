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

type technologyRepository interface {
	FindAll() ([]technology, error)
}

type technologyRepositoryImpl struct {
	db *sql.DB
}

func (r technologyRepositoryImpl) FindAll() ([]technology, error) {
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

func NewTechnologyRepository(db *sql.DB) technologyRepository {
	return technologyRepositoryImpl{db: db}
}
