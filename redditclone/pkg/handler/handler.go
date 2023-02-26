package handler

import (
	"github.com/gorilla/mux"
	"gitlab.com/vk-go/lectures-2022-2/06_databases/99_hw/redditclone/pkg/service"
	"net/http"
)

const (
	CategoryName = "CATEGORY_NAME"
	PostID       = "POST_ID"
	CommentID    = "COMMENT_ID"
	UserLogin    = "USER_LOGIN"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/", h.indexHTML)
	r.Handle("/static/", staticHandler)

	api := mux.NewRouter()
	r.Handle("/api/", api)

	api.HandleFunc("/api/register", h.register).Methods("POST")
	api.HandleFunc("/api/login", h.login).Methods("POST")
	api.HandleFunc("/api/posts/", h.getAllPosts).Methods("GET")
	api.HandleFunc("/api/posts/{"+CategoryName+"}", h.getPostsByCategory).Methods("GET")
	api.HandleFunc("/api/post/{"+PostID+"}", h.getPost).Methods("GET")
	api.HandleFunc("/api/user/{"+UserLogin+"}", h.getPostsByUserLogin).Methods("GET")

	authHandler := mux.NewRouter()
	authWithMiddlewareHandler := h.authMiddleware(authHandler)
	api.PathPrefix("/api/").Handler(authWithMiddlewareHandler)

	authHandler.HandleFunc("/api/posts", h.newPost).Methods("POST")
	authHandler.HandleFunc("/api/post/{"+PostID+"}", h.addComment).Methods("POST")
	authHandler.HandleFunc("/api/post/{"+PostID+"}/{"+CommentID+"}", h.deleteComment).Methods("DELETE")
	authHandler.HandleFunc("/api/post/{"+PostID+"}/upvote", h.upvote).Methods("GET")
	authHandler.HandleFunc("/api/post/{"+PostID+"}/downvote", h.downvote).Methods("GET")
	authHandler.HandleFunc("/api/post/{"+PostID+"}/unvote", h.unvote).Methods("GET")
	authHandler.HandleFunc("/api/post/{"+PostID+"}", h.deletePost).Methods("DELETE")

	return r
}
