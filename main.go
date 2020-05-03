package main

import (
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("environment variable PORT is empty!")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		logger.Info("header", zap.Any("headers", r.Header))
		switch r.Method {
		case http.MethodGet:
			logger.Info("get request", zap.String("url", r.URL.Path), zap.String("query", r.URL.Query().Encode()))
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("post request",
				zap.String("url", r.URL.Path),
				zap.String("query", r.URL.Query().Encode()),
				zap.String("body", string(body)),
			)
		default:
			logger.Info("undefined method type")
		}
		w.WriteHeader(http.StatusOK)
	})
	logger.Info("http-mock-server is listening.", zap.String("port", port))
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}
}
