package internal

type Command interface {
	Handle() (int, error)
	Validate() error
}
