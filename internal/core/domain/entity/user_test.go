package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	name := "Joneco"
	email := "joneco@joneco.com.br"
	password := "123456"
	role := user_role_enum_user
	user, err := NewUser(name, email, password, role)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)
	assert.NotEqual(t, role, user_role_enum_admin)
	assert.Equal(t, role, user_role_enum_user)

}

func TestUser_ValidatePassword(t *testing.T) {
	name := "Joneco"
	email := "joneco@joneco.com.br"
	password := "123456"
	role := user_role_enum_user
	user, err := NewUser(name, email, password, role)
	assert.Nil(t, err)
	assert.Nil(t, user.ValidatePassword("123456"))
	assert.Error(t, user.ValidatePassword("1234567"))
	// will garantee that password is hashed.
	assert.NotEqual(t, "123456", user.Password)
}
