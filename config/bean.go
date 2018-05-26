package config

// BeanScope type for BeanDefinition
type BeanScope string

// ScopeSingleton type for singleton beans
type ScopeSingleton BeanScope

// ScopePrototype type for prototype beans
type ScopePrototype BeanScope

// BeanDefinition type
type BeanDefinition struct {
	Name         string
	Scope        BeanScope
	Factory      func(ConfigurationContext) interface{}
	Dependencies []string
}

// ConfigurationContext have al beans definition for the application
type ConfigurationContext struct {
	BeanDefinitions []BeanDefinition
}
