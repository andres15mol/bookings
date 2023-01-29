package dbrepo

import (
	"database/sql"

	"github.com/andres15mol/bookings/internal/config"
	"github.com/andres15mol/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewTestingRepo( a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo {
		App: a,
	}
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB: conn,
	}
}