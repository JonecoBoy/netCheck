package user

import (
	"errors"
	"github.com/JonecoBoy/netCheck/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        entity.ID `json:"id"`
	Name      string
	Email     string
	Password  string `json:"-"`
	Role      int
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

const (
	user_role_enum_user = iota
	user_role_enum_admin
)

func (u *User) validateUser() error {
	// validate user Role
	switch u.Role {
	case user_role_enum_user, user_role_enum_admin:
		return nil
	default:
		return errors.New("invalid user role")
	}
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func NewUser(name string, email string, password string, role int) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &User{
		ID:        entity.NewId(),
		Name:      name,
		Email:     email,
		Password:  string(hash),
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	err = newUser.validateUser()
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
