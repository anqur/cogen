package cogen

import (
	"fmt"

	"github.com/anqur/cogen/internal"
)

type Cogen interface {
	fmt.Stringer
}

type Option interface {
	Apply(s *internal.Setting)
}

type tableName string

func (n tableName) Apply(s *internal.Setting) { s.TableName = string(n) }
func WithTableName(name string) Option        { return tableName(name) }

func newSetting(opts []Option) *internal.Setting {
	s := new(internal.Setting)
	for _, opt := range opts {
		opt.Apply(s)
	}
	return s
}

func MySQL(data interface{}, opts ...Option) (Cogen, error) {
	return internal.MySQL(data, newSetting(opts))
}
