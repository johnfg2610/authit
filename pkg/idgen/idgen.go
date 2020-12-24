package idgen

type IDGenerator interface {
	GeneratePassword() (string, error)
}
