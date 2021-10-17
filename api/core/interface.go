package core

type Constructable interface {
	Load() error
	Dump() string
}

type CompletionStatus interface {
	GetCompletionStatus() bool
}
