package entity

import (
	domainError "github.com/jonecoboy/netCheck/internal/core/domain/error"
	"github.com/jonecoboy/netCheck/pkg/entity"
	"time"
)

type Plan struct {
	ID          entity.ID    `json:"id" bson:"id"`
	Name        string       `json:"name" bson:"name"`
	Description string       `json:"description" bson:"description"`
	Price       entity.Money `json:"price" bson:"price"`
	CreatedAt   time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `bson:"updated_at" json:"updated_at"`
	DeletedAt   time.Time    `bson:"deleted_at" json:"deleted_at"`
}

func (p *Plan) validate() error {
	if p.Name == "" {
		return domainError.ErrInvalidPlanName
	}
	if !p.DeletedAt.IsZero() {
		return domainError.ErrPlanDeleted
	}
	return nil
}

func NewPlan(name string, description string, price entity.Money) (*Plan, error) {
	p := &Plan{
		ID:          entity.NewId(),
		Name:        name,
		Description: description,
		Price:       price,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	if err := p.validate(); err != nil {
		return nil, err
	}
	return p, nil
}
