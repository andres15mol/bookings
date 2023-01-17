package driver

import (
	"database/sql"
	"time"
	
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/jackc/pgx/v5"
)

//DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdLeBDConn = 5
const maxDbLifeTime = 5 * time.Minute

//ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)

	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdLeBDConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
//testDB tries to ping the database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}
//NewDataBase creates a new database for the application
func NewDataBase(dsn string) (*sql.DB, error){
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil{
		return nil, err
	}

	return db, nil
}