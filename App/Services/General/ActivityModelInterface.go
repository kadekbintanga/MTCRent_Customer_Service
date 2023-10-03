package General

type ActivityModelInterface interface {
	TableName() string
	SetProperty() map[string]interface{}
	SetReference() uint
}
