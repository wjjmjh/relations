package main

import (
	"net/http"
	"relations/logger"
	"relations/service"
	"strconv"

	"go.uber.org/zap"
)

var svcPort int

func init() {
	svcPort = 7000
}

func main() {
	r := service.NewRouter()
	http.Handle("/", r)

	logger.Logger.Info("Starting HTTP service at", zap.Int("SvcPort", svcPort))
	err := http.ListenAndServe(":"+strconv.Itoa(svcPort), nil)

	if err != nil {
		panic(err)
	}
}
