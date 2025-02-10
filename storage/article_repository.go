package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gweningwarr/petOne/internal/app/models"
	"log"
)

type ArticleRepository struct {
	storage *Storage
}

var (
	tableArticle = "articles"
)

func (ar *ArticleRepository) Create(article *models.Article) (*models.Article, error) {
	query := fmt.Sprintf("INSERT INTO %s (title, author, content) VALUES ($1, $2, $3) RETURNING id", tableArticle)

	err := ar.storage.db.QueryRow(query, article.Title, article.Author, article.Content).Scan(&article.ID)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return article, nil
}

func (ar *ArticleRepository) DeleteById(id int) (*models.Article, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableArticle)

	_, err := ar.storage.db.Exec(query, id)
	if err != nil {

		return nil, err
	}

	return nil, nil
}

func (ar *ArticleRepository) FindById(id int) (*models.Article, bool, error) {
	query := fmt.Sprintf("SELECT FROM %s WHERE id = $1", tableArticle)

	var article models.Article
	row := ar.storage.db.QueryRow(query, id)

	if err := row.Scan(&article.ID, &article.Title, &article.Author, &article.Content); errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Статья не найдена")
	} else if err != nil {
		log.Fatalf("Error: Unable to execute query")
	}

	return &article, true, nil
}

func (ar *ArticleRepository) SelectAll() ([]*models.Article, error) {

	query := fmt.Sprintf("SELECT * FROM %s", tableArticle)

	rows, err := ar.storage.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	articles := make([]*models.Article, 0, 100)

	for rows.Next() {
		a := models.Article{}

		if errScan := rows.Scan(&a.ID, &a.Title, &a.Author, &a.Content); errScan != nil {
			log.Println(errScan)
			continue
		}

		articles = append(articles, &a)
	}

	return articles, nil
}
