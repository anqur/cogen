package cogen

import (
	"fmt"

	"github.com/anqur/cogen/internal"
)

type Cogen interface {
	fmt.Stringer
}

func MySQL(data interface{}) (Cogen, error) {
	return internal.MySQL(data)
}
