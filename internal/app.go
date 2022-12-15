package internal

type App interface {
	Name() string
	Technology() string
	Version() string
	Path() string
	Create() error
}
