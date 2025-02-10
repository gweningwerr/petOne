package api

import (
	"github.com/gweningwarr/petOne/internal/app/middleware"
	"github.com/gweningwarr/petOne/storage"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	prefix string = "/api/v1"
)

func (api *API) configureLoggerField() error {
	logLevel, err := logrus.ParseLevel(api.config.LoggerLevel)

	if err != nil {
		return err
	}

	api.logger.SetLevel(logLevel)

	return nil
}

func (api *API) configureRouterField() {

	api.router.HandleFunc(prefix+"/articles", api.PostArticle).Methods("POST")
	api.router.HandleFunc(prefix+"/articles", api.GetAllArticles).Methods("GET")
	api.router.Handle(
		prefix+"/article/{id}",
		middleware.JwtMiddleware.Handler(
			http.HandlerFunc(api.GetArticleById),
		)).Methods("GET")
	api.router.HandleFunc(prefix+"/article/{id}", api.DeleteArticleById).Methods("DELETE")

	api.router.HandleFunc(prefix+"/user", api.PostUser).Methods("POST")
	api.router.HandleFunc(prefix+"/users", api.GetUsers).Methods("GET")
	api.router.HandleFunc(prefix+"/user/{login}", api.GetUser).Methods("GET")

	api.router.HandleFunc(prefix+"/user/auth", api.PostToAuth).Methods("POST")

	api.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			return
		}
	})
}

func (api *API) configureStorageField() error {
	storage := storage.New(api.config.Storage)

	if err := storage.Open(); err != nil {
		return err
	}

	api.storage = storage

	return nil
}
