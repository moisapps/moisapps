package infrastructure

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

const InitDb string = `
	CREATE TABLE IF NOT EXISTS TECHNOLOGY(
	    NAME TEXT NOT NULL PRIMARY KEY
	);
	CREATE TABLE IF NOT EXISTS VERSION(
	    NUMBER TEXT NOT NULL,
	    TECHNOLOGY_NAME TEXT NOT NULL,
	    PRIMARY KEY(NUMBER, TECHNOLOGY_NAME),
	    FOREIGN KEY(TECHNOLOGY_NAME) REFERENCES TECHNOLOGY(NAME) 
	)
`

const InitDDL string = `
	INSERT INTO TECHNOLOGY VALUES ('JAVA');
	INSERT INTO TECHNOLOGY VALUES ('NODE');
	INSERT INTO TECHNOLOGY VALUES ('GOLANG');
	INSERT INTO VERSION VALUES ('8', 'JAVA');
	INSERT INTO VERSION VALUES ('11', 'JAVA');
	INSERT INTO VERSION VALUES ('17', 'JAVA');
	INSERT INTO VERSION VALUES ('18.12.1', 'NODE');
	INSERT INTO VERSION VALUES ('16.17.1', 'NODE');
	INSERT INTO VERSION VALUES ('14.4.0', 'NODE');
	INSERT INTO VERSION VALUES ('19', 'GOLANG');
	INSERT INTO VERSION VALUES ('18', 'GOLANG');
	INSERT INTO VERSION VALUES ('17', 'GOLANG');
`

const filename string = "/usr/local/etc/moisapps/mspdb.dbd"

var DB *sql.DB

func SetupDatabase() {
	driverName := "sqlite3"
	connectionString := filename
	dbTemp, err := sql.Open(driverName, connectionString)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		if _, err = dbTemp.Exec(InitDb); err != nil {
			panic(err)
		}
		if _, err = dbTemp.Exec(InitDDL); err != nil {
			panic(err)
		}
	}

	DB = dbTemp
}
