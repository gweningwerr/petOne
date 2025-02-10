package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gweningwarr/petOne/internal/app/models"
	"log"
)

type UserRepository struct {
	storage *Storage
}

var (
	tableUser string = "users"
)

func (ur *UserRepository) Create(u *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING id", tableUser)

	err := ur.storage.db.QueryRow(query, u.Login, u.Password).Scan(&u.ID)

	if err != nil {
		log.Println(query)
		return nil, err
	}

	return u, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	query := fmt.Sprintf("SELECT id, login, password, phone_number FROM %s WHERE login = $1", tableUser)

	var user models.User
	row := ur.storage.db.QueryRow(query, login)

	if err := row.Scan(&user.ID, &user.Login, &user.Password, &user.PhoneNumber); errors.Is(err, sql.ErrNoRows) {
		log.Println("Пользователь не найден")
		return nil, false, nil
	} else if err != nil {
		log.Fatalf("Error: Unable to execute query, %v", err)
		return nil, false, err
	}

	return &user, true, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {

	query := fmt.Sprintf("SELECT * FROM %s", tableUser)

	rows, err := ur.storage.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]*models.User, 0, 100)

	for rows.Next() {
		u := models.User{}

		if errScan := rows.Scan(&u.ID, &u.Login, &u.Password, &u.PhoneNumber); errScan != nil {
			log.Println(errScan)
			continue
		}

		users = append(users, &u)
	}
	return users, nil
}
