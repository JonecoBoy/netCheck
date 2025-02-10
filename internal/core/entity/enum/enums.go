package enum

// Common Core enums used to multiple packages
const (
	Status_enum_created = iota
	Status_enum_pending
	Status_enum_approved
	Status_enum_rejected
)

const (
	Role_enum_view      = 1
	Role_enum_edit      = 2 << 0
	Role_enum_delete    = 2 << 1
	Role_enum_manage    = 2 << 2
	Role_enum_financial = 2 << 3
)
