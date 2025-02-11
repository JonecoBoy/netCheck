package entity

import (
	entityPlan "github.com/jonecoboy/netCheck/internal/core/domain/entity/plan"
	entityUser "github.com/jonecoboy/netCheck/internal/core/domain/entity/user"
	domainError "github.com/jonecoboy/netCheck/internal/core/domain/error"
	pkgEntity "github.com/jonecoboy/netCheck/pkg/entity"
)

type Subscription struct {
	ID   pkgEntity.ID
	User entityUser.User
	Plan entityPlan.Plan
}

func (s *Subscription) validate() error {
	if !s.User.DeletedAt.IsZero() {
		return domainError.ErrUserDeleted
	}
	return nil
}

func NewSubscription(user *entityUser.User, plan *entityPlan.Plan) (*Subscription, error) {
	s := &Subscription{
		ID:   pkgEntity.NewId(),
		User: *user,
		Plan: *plan,
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return s, nil
}
