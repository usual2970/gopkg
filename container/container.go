package container

import (
	"sync"

	"go.uber.org/dig"
)

var c *dig.Container

var cOnce sync.Once

func GetContainer() *dig.Container {
	cOnce.Do(func() {
		c = dig.New()
	})
	return c
}

func Provide(constructor interface{}, options ...dig.ProvideOption) error {
	c := GetContainer()
	return c.Provide(constructor, options...)
}

func Invode(function interface{}, opts ...dig.InvokeOption) error {
	c := GetContainer()
	return c.Invoke(function, opts...)
}
