package types

type Type int

const (
	NoOp = iota + 1
	Zap
)

func Atot(a string) Type {
	switch a {
	case "zap":
		return Zap
	default:
		return NoOp
	}
}
