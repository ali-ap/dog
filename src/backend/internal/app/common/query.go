package internal

type Query interface {
	Handle() (interface{}, error)
	Validate() error
}
