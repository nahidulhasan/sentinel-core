package usercase

import (
    "errors"

    "github.com/nahidulhasan/sentinel-core/internal/domain/user"
)

type UserService interface {
    Create(name, email string) (*user.User, error)
    List() ([]*user.User, error)
}

type userService struct {
    repo user.Repository
}

func NewUserService(r user.Repository) UserService {
    return &userService{repo: r}
}

func (s *userService) Create(name, email string) (*user.User, error) {
    if name == "" || email == "" {
        return nil, errors.New("name and email required")
    }
    u := &user.User{
        Name:  name,
        Email: email,
    }
    if err := s.repo.Create(u); err != nil {
        return nil, err
    }
    return u, nil
}

func (s *userService) List() ([]*user.User, error) {
    return s.repo.List()
}
