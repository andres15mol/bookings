package dbrepo

import (
	"database/sql"

	"github.com/andres15mol/bookings/internal/config"
	"github.com/andres15mol/bookings/internal/models"
	"github.com/andres15mol/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// InsertResevation implements repository.DatabaseRepo
func (*testDBRepo) InsertResevation(res models.Reservation) (int, error) {
	panic("unimplemented")
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
		DB:  &sql.DB{},
	}
}
