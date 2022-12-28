package entity

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const technologyName = "NODE"

type nodeApp struct {
	name             string
	version          string
	path             string
	createAppUseCase createNodeApp
	repository       nodeRepository
}

func NewNodeApp(name, version, path string, db *sql.DB) *nodeApp {
	return &nodeApp{
		name:             name,
		version:          version,
		path:             path,
		createAppUseCase: newCreateNodeApp(db, name, path, version),
		repository:       newNodeRepository(db),
	}
}

func (n nodeApp) Name() string {
	return n.name
}

func (n nodeApp) Technology() string {
	return technologyName
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

type nodeRepository interface {
	findSupportedVersions() ([]string, error)
}

type nodeRepositoryImpl struct {
	db *sql.DB
}

func (r nodeRepositoryImpl) findSupportedVersions() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutInMillis)*time.Millisecond)
	defer cancel()
	rows, err := r.db.QueryContext(ctx, "SELECT V.NUMBER FROM VERSION V WHERE V.TECHNOLOGY_NAME = $1 ORDER BY V.NUMBER ASC", technologyName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var versions []string
	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	return versions, nil
}

func newNodeRepository(db *sql.DB) nodeRepository {
	return nodeRepositoryImpl{db: db}
}

type createNodeApp interface {
	execute(n nodeApp) error
}

type osCommand interface {
	Run() error
}

type createNodeAppUseCase struct {
	db       *sql.DB
	out      *bytes.Buffer
	commands []osCommand
}

func (useCase createNodeAppUseCase) execute(n nodeApp) error {
	versions, err := n.repository.findSupportedVersions()
	if err != nil {
		return err
	}
	var supportedVersion bool
	for _, version := range versions {
		if strings.Compare(strings.TrimSpace(n.version), strings.TrimSpace(version)) == 0 {
			supportedVersion = true
		}
	}
	if !supportedVersion {
		return errors.New("version not supported")
	}
	for _, cmd := range useCase.commands {
		err = cmd.Run()
		if err != nil {
			fmt.Println(useCase.out.String())
			return err
		}
	}
	fmt.Println(strings.TrimSpace(useCase.out.String()))
	return nil
}

func newCreateNodeApp(db *sql.DB, name, path, version string) createNodeApp {
	var out bytes.Buffer
	var commands []osCommand
	commands = append(commands, changeNodeVersionCmd(version, &out))
	if strings.Compare(".", strings.TrimSpace(path)) != 0 {
		projectDirectoryCmd, projectPath := createProjectDirectoryCmd(name, path)
		commands = append(commands, projectDirectoryCmd)
		path = projectPath
	}
	commands = append(commands, npmInitCmd(path, &out))
	return createNodeAppUseCase{
		db:       db,
		commands: commands,
		out:      &out,
	}
}

func changeNodeVersionCmd(version string, out *bytes.Buffer) *exec.Cmd {
	command := exec.Command("bash", "-c", fmt.Sprintf("source ~/.nvm/nvm.sh && nvm use %s", version))
	command.Stdout = out
	return command
}

func createProjectDirectoryCmd(name string, path string) (*exec.Cmd, string) {
	if len(strings.TrimSpace(path)) < 1 {
		path = "/tmp/moisapps"
	}
	path = fmt.Sprintf("%s/%s", path, name)
	command := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s", path))
	return command, path
}

func npmInitCmd(path string, out *bytes.Buffer) *exec.Cmd {
	command := exec.Command("bash", "-c", fmt.Sprintf("cd %s && npm init -y", path))
	command.Stdout = out
	return command
}
