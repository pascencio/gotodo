package config

// BeanScope type for BeanDefinition
type BeanScope string

// ScopeSingleton type for singleton beans
const ScopeSingleton BeanScope = "singleton"

// ScopePrototype type for prototype beans
const ScopePrototype BeanScope = "prototype"

// BeanDefinition type
type BeanDefinition struct {
	Name         string
	Scope        BeanScope
	Factory      func(ConfigurationContext) interface{}
	Dependencies []string
}

// GetBean return a intance of bean definition calling the factory method
func (b BeanDefinition) GetBean(c ConfigurationContext) interface{} {
	return b.Factory(c)
}

// ConfigurationContext have al beans definition for the application
type ConfigurationContext struct {
	BeanDefinitions map[string]BeanDefinition
}
