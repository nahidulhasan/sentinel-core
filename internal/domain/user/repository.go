package user

type Repository interface {
    Create(u *User) error
    List() ([]*User, error)
}
