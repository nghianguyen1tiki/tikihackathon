package logger

type Mode int

const (
	Production = iota + 1
	Development
)

func Atom(a string) Mode {
	switch a {
	case "production":
		return Production
	case "development":
		return Development
	default:
		return Development
	}
}
