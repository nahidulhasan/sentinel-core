package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/jmoiron/sqlx"
    "github.com/go-chi/chi/v5"
    "github.com/nahidulhasan/sentinel-core/internal/domain/user"
    "github.com/nahidulhasan/sentinel-core/internal/usecase/usercase"
)


func NewUserRouter(db *sqlx.DB) http.Handler {
    r := chi.NewRouter()

    repo := user.NewPostgresRepo(db)
    svc := usercase.NewUserService(repo)

    r.Get("/users", ListUsers(svc))
    r.Post("/users", CreateUser(svc))

    return r
}

func ListUsers(svc usercase.UserService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        users, err := svc.List()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(users)
    }
}

func CreateUser(svc usercase.UserService) http.HandlerFunc {
    type req struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    return func(w http.ResponseWriter, r *http.Request) {
        var body req
        if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
            http.Error(w, "invalid body", http.StatusBadRequest)
            return
        }
        u, err := svc.Create(body.Name, body.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(u)
    }
}
