package service

import (
	"encoding/json"
	"net/http"
	"relations/logger"

	"go.uber.org/zap"
)

func OkResponse(w http.ResponseWriter, res interface{}) {
	resBytes, err := json.Marshal(res)
	if err != nil {
		logger.Logger.Error("OkResponse: json.Marshal", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resBytes)
	if err != nil {
		logger.Logger.Error("OkResponse: w.Write(res): %v", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
