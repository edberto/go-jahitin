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

	OrderStatusAtoI = map[string]int{
		OrderStatusWaitingA:    OrderStatusWaitingI,
		OrderStatusRejectedA:   OrderStatusRejectedI,
		OrderStatusProcessingA: OrderStatusProcessingI,
		OrderStatusInDeliveryA: OrderStatusInDeliveryI,
		OrderStatusDoneA:       OrderStatusDoneI,
	}

	OrderStatusItoA = map[int]string{
		OrderStatusWaitingI:    OrderStatusWaitingA,
		OrderStatusRejectedI:   OrderStatusRejectedA,
		OrderStatusProcessingI: OrderStatusProcessingA,
		OrderStatusInDeliveryI: OrderStatusInDeliveryA,
		OrderStatusDoneI:       OrderStatusDoneA,
	}
)

const (
	CustomerRoleA          = "customer"
	CustomerRoleI          = 1
	TailorRoleA            = "tailor"
	TailorRoleI            = 2
	OrderStatusWaitingA    = "waiting"
	OrderStatusWaitingI    = 1
	OrderStatusRejectedA   = "rejected"
	OrderStatusRejectedI   = 2
	OrderStatusProcessingA = "processing"
	OrderStatusProcessingI = 3
	OrderStatusInDeliveryA = "in-delivery"
	OrderStatusInDeliveryI = 4
	OrderStatusDoneA       = "done"
	OrderStatusDoneI       = 5
)
