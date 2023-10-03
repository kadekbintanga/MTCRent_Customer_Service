package Server

type ActivityAction struct{}

const ACTION_CREATE = "create"
const ACTION_UPDATE = "update"
const ACTION_DELETE = "delete"
const ACTION_GENERAL = "general"

func (srv ActivityAction) OptionCodes() []string {
	return []string{
		ACTION_CREATE,
		ACTION_UPDATE,
		ACTION_DELETE,
		ACTION_GENERAL,
	}
}
