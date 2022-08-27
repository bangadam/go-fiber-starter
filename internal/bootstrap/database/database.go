package database

import (
	"fmt"
	"log"

	"github.com/bangadam/go-fiber-starter/app/database/schema"
	"github.com/bangadam/go-fiber-starter/app/database/seeds"
	"github.com/bangadam/go-fiber-starter/utils/config"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setup database with gorm
type Database struct {
	DB  *gorm.DB
	Log zerolog.Logger
	Cfg *config.Config
}

func NewDatabase(cfg *config.Config, log zerolog.Logger) *Database {
	db := &Database{
		Cfg: cfg,
		Log: log,
	}

	return db
}

// connect database
func (_db *Database) ConnectDatabase() {
	conn, err := gorm.Open(postgres.Open(_db.Cfg.DB.Postgres.DSN), &gorm.Config{})
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to connect the database!")
	} else {
		_db.Log.Info().Msg("Connected the database succesfully!")
	}

	_db.DB = conn
}

// shutdown database
func (_db *Database) ShutdownDatabase() {
	sqlDB, err := _db.DB.DB()
	if err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to shutdown the database!")
	} else {
		_db.Log.Info().Msg("Shutdown the database succesfully!")
	}
	sqlDB.Close()
}

// migrate models
func (_db *Database) MigrateModels() {
	fmt.Println("----------------------------------")
	fmt.Println("| Running Migration:  |")
	fmt.Println("----------------------------------")
	if err := _db.DB.AutoMigrate(
		Models()...,
	); err != nil {
		_db.Log.Error().Err(err).Msg("An unknown error occurred when to migrate the database!")
	}
}

// list of models for migration
func Models() []interface{} {
	return []interface{}{
		schema.Article{},
	}
}

// seed data
func (_db *Database) SeedData() {
	for _, seed := range seeds.All() {
		fmt.Println("----------------------------------")
		fmt.Println("| Running Seed: " + seed.Name + " |")
		fmt.Println("----------------------------------")
		if err := seed.Run(_db.DB); err != nil {
			log.Fatal(err)
		}
	}
}
