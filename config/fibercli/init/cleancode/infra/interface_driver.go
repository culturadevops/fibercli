package infra

type Driver interface {
	NewDriver() error
	GetClient() interface{}
}
