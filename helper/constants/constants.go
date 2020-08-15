package constants

var (
	RoleAtoI = map[string]int{
		CustomerRoleA: CustomerRoleI,
		TailorRoleA:   TailorRoleI,
	}

	RoleItoA = map[int]string{
		CustomerRoleI: CustomerRoleA,
		TailorRoleI:   TailorRoleA,
	}
)

const (
	CustomerRoleA = "customer"
	CustomerRoleI = 1
	TailorRoleA   = "tailor"
	TailorRoleI   = 2
)
