package user_test

import (
	"github.com/golang/mock/gomock"
	"github.com/jonecoboy/netCheck/internal/core/application/usecase/user"
	"github.com/jonecoboy/netCheck/internal/core/domain/entity"
	"github.com/jonecoboy/netCheck/internal/core/domain/repository/mocks"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	userEntity := &entity.User{
		ID:        pkgEntity.NewId(),
		Name:      "Alice",
		Email:     "alice@example.com",
		Role:      entity.UserRoleEnumUser,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.EXPECT().
		Create(userEntity).
		Return(nil).
		Times(1)

	err := user.CreateUser(mockRepo, userEntity)
	assert.NoError(t, err)
}

func TestCreateUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	userEntity := &entity.User{
		ID:    pkgEntity.NewId(),
		Name:  "Bob",
		Email: "bob@example.com",
		Role:  entity.UserRoleEnumUser,
	}

	mockRepo.EXPECT().
		Create(userEntity).
		Return(assert.AnError).
		Times(1)

	err := user.CreateUser(mockRepo, userEntity)
	assert.Error(t, err)
}
