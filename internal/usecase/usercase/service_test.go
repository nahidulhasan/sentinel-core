package usercase

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/nahidulhasan/sentinel-core/internal/domain/user"
)

// mock repo
type mockRepo struct {
    users []*user.User
    err   error
}

func (m *mockRepo) Create(u *user.User) error {
    if m.err != nil {
        return m.err
    }
    u.ID = int64(len(m.users) + 1)
    m.users = append(m.users, u)
    return nil
}

func (m *mockRepo) List() ([]*user.User, error) {
    if m.err != nil {
        return nil, m.err
    }
    return m.users, nil
}

func TestCreateAndList(t *testing.T) {
    mr := &mockRepo{}
    svc := NewUserService(mr)

    u, err := svc.Create("Alice", "alice@example.com")
    assert.NoError(t, err)
    assert.Equal(t, int64(1), u.ID)

    list, err := svc.List()
    assert.NoError(t, err)
    assert.Len(t, list, 1)
}

func TestCreateValidation(t *testing.T) {
    mr := &mockRepo{}
    svc := NewUserService(mr)

    _, err := svc.Create("", "")
    assert.Error(t, err)
    assert.Equal(t, "name and email required", err.Error())
}

func TestRepoError(t *testing.T) {
    mr := &mockRepo{err: errors.New("db fail")}
    svc := NewUserService(mr)

    _, err := svc.Create("Bob", "bob@example.com")
    assert.Error(t, err)
}
