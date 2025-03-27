package entity

import (
	"github.com/jonecoboy/netCheck/internal/utils"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"go.mongodb.org/mongo-driver/mongo/address"
	"math/rand"
	"time"
)

var names = []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank"}
var domains = []string{"example.com", "test.com", "mail.net", "domain.org"}

// RandomUser returns a random but valid (non-deleted) User.
func RandomUser() *User {
	now := time.Now()
	name := names[rand.Intn(len(names))]
	email := generateEmail(name)

	return &User{
		ID:        pkgEntity.NewId(),
		Name:      name,
		Email:     email,
		Password:  "", // Omitted
		Role:      UserRoleEnumUser,
		Addresses: []address.Address{}, // Optional: Add dummy address if needed
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{}, // Not deleted
	}
}

// DeletedUser returns a user with DeletedAt set
func DeletedUser() *User {
	user := RandomUser()
	user.DeletedAt = time.Now()
	return user
}

// RandomAdmin returns a valid admin user
func RandomAdmin() *User {
	user := RandomUser()
	user.Role = UserRoleEnumAdmin
	return user
}

// generateEmail returns an email like "alice123@example.com"
func generateEmail(name string) string {
	return name + utils.RandomDigits(3) + "@" + domains[rand.Intn(len(domains))]
}
