package config

type BeanScope string

type ScopeSingleton BeanScope

type ScopePrototype BeanScope

type BeanDefinition struct {
	Name         string
	Scope        BeanScope
	Factory      func(ConfigurationContext) interface{}
	Dependencies []string
}

type ConfigurationContext struct {
	BeanDefinitions []BeanDefinition
}
