package error

import "errors"

var (
	ErrInvalidUser     = errors.New("Invalid user")
	ErrUserDeleted     = errors.New("User is deleted")
	ErrInvalidPlanName = errors.New("Invalid plan name")
	ErrPlanDeleted     = errors.New("Plan cannot be Deleted")
)
