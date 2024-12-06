package api

import (
	"github.com/gweningwarr/petOne/storage"
	"github.com/sirupsen/logrus"
	"net/http"
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
