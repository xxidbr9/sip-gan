package uuid

type ID string
type IDGenerator interface {
	GenerateID() (ID, error)
}