package scp

type ISCPDatabase interface {
	Connect() error
	GetObject(id string) (ISCPObject, error)
}
