package entity

import (
	"bytes"
	"database/sql"
	"fmt"
	"os/exec"
	"strings"
)

type nodeApp struct {
	name             string
	version          string
	path             string
	createAppUseCase createNodeApp
}

func NewNodeApp(name, version, path string, db *sql.DB) *nodeApp {
	return &nodeApp{
		name:             name,
		version:          version,
		path:             path,
		createAppUseCase: newCreateNodeApp(db),
	}
}

func (n nodeApp) Name() string {
	return n.name
}

func (n nodeApp) Technology() string {
	return "Node.js"
}

func (n nodeApp) Version() string {
	return n.version
}

func (n nodeApp) Path() string {
	return n.path
}

func (n nodeApp) Create() error {
	return n.createAppUseCase.execute(n)
}

type createNodeApp interface {
	execute(n nodeApp) error
}

type createNodeAppUseCase struct {
	db *sql.DB
}

func (useCase createNodeAppUseCase) execute(n nodeApp) error {
	fmt.Printf("%s\n%s\n%s\n", n.name, n.version, n.path)
	var out bytes.Buffer
	execCmd := exec.Command("npm", "--version")
	execCmd.Stdout = &out
	err := execCmd.Run()
	if err != nil {
		return err
	}
	fmt.Println(strings.TrimSpace(out.String()))
	return nil
}

func newCreateNodeApp(db *sql.DB) createNodeApp {
	return createNodeAppUseCase{db: db}
}
