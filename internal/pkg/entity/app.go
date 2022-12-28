package entity

import (
	"database/sql"
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

func NewApplication(name, tech, version, path string) application {
	if strings.Contains(strings.ToUpper(tech), "NODE") {
		tech = "NODE"
	}
	return application{
		name:    name,
		tech:    tech,
		version: version,
		path:    path,
	}
}

func (a application) Create(db *sql.DB) error {
	var err error
	switch strings.ToUpper(a.tech) {
	case "NODE":
		nodeApp := NewNodeApp(a.name, a.version, a.path, db)
		err = nodeApp.Create()
	}
	return err
}
