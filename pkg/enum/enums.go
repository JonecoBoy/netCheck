package enum

// Common Core enums used to multiple packages

type StatusEnum int

type RoleEnum int

const (
	_ StatusEnum = iota
	Created
	Pending
	Approved
	Rejected
)

const (
	View RoleEnum = 1 << iota
	Write
	Edit
	Delete
)
