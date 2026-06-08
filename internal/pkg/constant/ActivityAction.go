package constant

const ACTION_CREATE = "create"
const ACTION_UPDATE = "update"
const ACTION_DELETE = "delete"
const ACTION_GENERAL = "general"
const ACTION_STATUS = "status"

/** --- SUB FEATURE --- */
const ACTIVITY_CUSTOMER_UPDATE_STATUS = "customer_update_status"

type ActivityAction struct{}

func (srv ActivityAction) OptionCodeNames() []string {
	return []string{
		ACTION_CREATE,
		ACTION_UPDATE,
		ACTION_DELETE,
		ACTION_GENERAL,
	}
}
