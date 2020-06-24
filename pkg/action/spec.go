package action

type Interface interface {
	Execute() error
	Explain() string
}
