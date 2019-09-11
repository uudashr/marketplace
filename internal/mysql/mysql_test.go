// +build integration

package mysql_test

import (
	"database/sql"
	"flag"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	dbScripts  = flag.String("db-scripts", "file://migrations", "The location of migration scripts.")
	dbUser     = flag.String("db-user", "marketplace", "Database username")
	dbPassword = flag.String("db-password", "secret", "Database password")
	dbAddress  = flag.String("db-address", "localhost:3306", "Database address")
	dbName     = flag.String("db-name", "marketplace_test", "Database name")
)

const driverName = "mysql"

type dbFixture struct {
	t  *testing.T
	m  *migrate.Migrate
	db *sql.DB
}

func (s *dbFixture) tearDown() {
	if err := s.m.Down(); err != nil {
		if err != migrate.ErrNoChange {
			s.t.Error("Fail to execute migration down scripts", err)
		}
	}

	if err := s.db.Close(); err != nil {
		s.t.Error("Fail to close db:", err)
	}
}

func setupDBFixture(t *testing.T) *dbFixture {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true&clientFoundRows=true&parseTime=true&loc=Local", *dbUser, *dbPassword, *dbAddress, *dbName)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		t.Fatal("Fail to open DB:", err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		t.Fatal("Fail to create driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(*dbScripts, driverName, driver)
	if err != nil {
		t.Fatal("Fail to create DB instance:", err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			t.Error("Fail to execute migration up scripts:", err)
		}
	}

	return &dbFixture{
		t:  t,
		m:  m,
		db: db,
	}
}
