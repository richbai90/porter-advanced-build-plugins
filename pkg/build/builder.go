package build

import (
	"github.com/richbai90/porter-advanced-build-plugins/pkg/config"
)

type Builder struct {
	config config.Config
}

func NewBuilder(cfg config.Config) *Builder {
	b := &Builder{
		config: cfg,
	}

	return b
}

func (b *Builder) Command() string {
	return "TODO"
}
