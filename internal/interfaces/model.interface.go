package interfaces

type IModel interface {
	GetID() uint
	SetID(uint)
	TableName() string
}
