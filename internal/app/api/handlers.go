package api

import (
	"encoding/json"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
	"github.com/gweningwarr/petOne/internal/app/middleware"
	"github.com/gweningwarr/petOne/internal/app/models"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHandlers(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Getting all articles")

	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error getting all articles: ", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error getting all articles",
			IsError:    true,
		}

		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(articles)
}

func (api *API) PostArticle(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Postting new article")

	var article models.Article
	err := json.NewDecoder(request.Body).Decode(&article)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			IsError:    true,
		}
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Error creating article (DB connect): ", err)
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error creating article",
			IsError:    true,
		}
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(a)
}

func (api *API) GetArticleById(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Getting article by ID")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Error getting article by ID: ", err)

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Error getting article by ID",
			IsError:    true,
		}

		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)
	}

	article, finded, err := api.storage.Article().FindById(id)

	if err != nil {
		api.logger.Info("Error getting article by ID: ", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error getting article by ID",
			IsError:    true,
		}

		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	if !finded {
		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "Article not found",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(article)
}

func (api *API) DeleteArticleById(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Deleting article by ID")

	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		api.logger.Info("Error delete article by ID: ", err)
	}

	r, err := api.storage.Article().DeleteById(id)
	if err != nil {
		api.logger.Info("Error deleting article by ID: ", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error deleting article by ID",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
	}

	writer.WriteHeader(http.StatusNoContent)
	json.NewEncoder(writer).Encode(r)
}

func (api *API) PostUser(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Postting new user")

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			IsError:    true,
		}

		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	u, err := api.storage.User().Create(&user)
	if err != nil {
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error creating user",
			IsError:    true,
		}

		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)

		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(u)
}

func (api *API) GetUsers(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Getting all users")

	users, err := api.storage.User().SelectAll()
	if err != nil {
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error getting all users",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(users)
}

func (api *API) GetUser(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Getting user")

	login, exist := mux.Vars(request)["login"]

	if !exist {
		api.logger.Info("Не удалось получить логин для поиска: ", exist)
		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    "Не удалось получить логин для поиска",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	user, founded, errUser := api.storage.User().FindByLogin(login)

	if errUser != nil {
		api.logger.Info("Error getting user ID: ", errUser)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error getting user ID",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !founded {
		api.logger.Info("User not found")

		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(user)
}

func (api *API) PostToAuth(writer http.ResponseWriter, request *http.Request) {
	initHandlers(writer)

	api.logger.Info("Post to auth api/v1/user/auth")

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		api.logger.Info("Не корректный json ", err)

		msg := Message{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	userInDB, ok, err := api.storage.User().FindByLogin(user.Login)

	if err != nil {
		api.logger.Info("Error getting user ID: ", err)
		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error getting user ID",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if !ok {
		api.logger.Info("User with that login does not exists")
		msg := Message{
			StatusCode: http.StatusNotFound,
			Message:    "User not found",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	if userInDB.Password != user.Password {
		api.logger.Info("Invalid password")
		msg := Message{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid password",
			IsError:    true,
		}

		writer.WriteHeader(msg.StatusCode)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // дополнительные действия для шифрования
	claims["id"] = user.ID
	claims["login"] = user.Login
	claims["exp"] = time.Now().Add(time.Hour / 6).Unix()

	tokenString, err := token.SignedString(middleware.SecretKey)

	if err != nil {
		api.logger.Info("Error signing token: ", err)

		msg := Message{
			StatusCode: http.StatusInternalServerError,
			Message:    "We have some trouble. Try again later.",
			IsError:    true,
		}

		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: http.StatusOK,
		Message:    tokenString,
		IsError:    false,
	}

	writer.WriteHeader(msg.StatusCode)
	json.NewEncoder(writer).Encode(msg)
}
