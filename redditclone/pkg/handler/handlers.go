package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

var staticHandler = http.StripPrefix(
	"/static/",
	http.FileServer(http.Dir("./static")),
)

func (h *Handler) indexHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/html/index.html")
}

func JSONResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	logrus.Println(message)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	logrus.Println(err)
}
