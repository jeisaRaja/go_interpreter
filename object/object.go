package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

type Object interface {
	Inspect() string
	Type() ObjectType
}

type Integer struct {
	Value int64
}

func (in *Integer) Inspect() string {
	return fmt.Sprintf("%d", in.Value)
}
func (in *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
