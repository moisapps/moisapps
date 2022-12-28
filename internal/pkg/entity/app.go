package entity

import (
	"database/sql"
	"errors"
	"strings"
)

type App interface {
	Name() string
	Technology() string
	Version() string
	Path() string
	Create() error
}

type application struct {
	name    string
	tech    string
	version string
	path    string
}

func NewApplication(name, tech, version, path string) (application, error) {
	if len(strings.TrimSpace(name)) < 1 {
		return application{}, errors.New("informar o nome da aplicação é obrigatório")
	}
	if len(strings.TrimSpace(tech)) < 1 {
		return application{}, errors.New("informar o nome da tecnologia é obrigatório")
	}
	if len(strings.TrimSpace(version)) < 1 {
		return application{}, errors.New("informar a versão da tecnologia é obrigatório")
	}
	if strings.Contains(strings.ToUpper(tech), "NODE") {
		tech = "NODE"
	}
	return application{
		name:    name,
		tech:    tech,
		version: version,
		path:    path,
	}, nil
}

func (a application) Create(db *sql.DB) error {
	var err error
	switch strings.ToUpper(a.tech) {
	case "NODE":
		nodeApp := NewNodeApp(a.name, a.version, a.path, db)
		err = nodeApp.Create()
	default:
		err = errors.New("technologia não suportada ou não implementada até o momento")
	}
	return err
}
