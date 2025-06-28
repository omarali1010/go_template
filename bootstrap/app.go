package bootstrap

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/omaraliali1010/go_template/internal/db"
	"github.com/omaraliali1010/go_template/internal/jwtservice"
)

type Application struct {
	Env        *Env
	DB         *sql.DB
	Queries    *db.Queries
	JWTService *jwtservice.JWTService
}

func App() *Application {
	app := &Application{}
	var err error

	app.Env, err = NewEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Construct PostgreSQL connection string
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		app.Env.DBHost,
		app.Env.DBPort,
		app.Env.DBUser,
		app.Env.DBPass,
		app.Env.DBName,
	)

	// Open DB connection
	app.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	// Configure connection pool
	app.DB.SetMaxOpenConns(25)
	app.DB.SetMaxIdleConns(25)
	app.DB.SetConnMaxLifetime(5 * time.Minute)

	// Verify connection
	if err = app.DB.Ping(); err != nil {
		log.Fatal("failed to ping db:", err)
	}

	// Initialize sqlc Queries struct
	app.Queries = db.New(app.DB)

	app.JWTService = &jwtservice.JWTService{
		AccessSecret:       app.Env.AccessTokenSecret,
		RefreshSecret:      app.Env.RefreshTokenSecret,
		AccessTokenExpiry:  time.Duration(app.Env.AccessTokenExpiryHour) * time.Hour,
		RefreshTokenExpiry: time.Duration(app.Env.RefreshTokenExpiryHour) * time.Hour,
	}

	return app
}

func (app *Application) CloseDBConnection() {
	if err := app.DB.Close(); err != nil {
		log.Println("error closing db:", err)
	}
}
