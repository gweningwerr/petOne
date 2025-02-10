package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // Для того чтобы отработала функция init()
	"log"
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	log.Printf("Opening database: %s", storage.config.DatabaseURI)
	db, err := sql.Open("postgres", storage.config.DatabaseURI)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	storage.db = db
	log.Println("Connected to database")

	return nil
}

func (storage *Storage) Close() {
	err := storage.db.Close()
	if err != nil {
		return
	}
}

func (storage *Storage) User() *UserRepository {
	if storage.userRepository == nil {
		storage.userRepository = &UserRepository{
			storage: storage,
		}
	}

	return storage.userRepository
}

func (storage *Storage) Article() *ArticleRepository {
	if storage.articleRepository == nil {
		storage.articleRepository = &ArticleRepository{
			storage: storage,
		}
	}

	return storage.articleRepository
}
