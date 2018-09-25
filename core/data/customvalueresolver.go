package data

import "sync"

var (
 resolvers []CustomValueResolver
 lock = &sync.Mutex{}
)


type CustomValueResolver interface {
	CanResolveValue(toResolve string) bool
	ResolveValue(toResolve string, scope Scope) (interface{}, error)
}


func RegisterCustomValueResolver(resolver CustomValueResolver) {
	lock.Lock()
	defer lock.Unlock()
	resolvers = append(resolvers, resolver)
}


func GetCustomValueResolver(toResolve string) CustomValueResolver  {
	for _, resolver := range resolvers {
		if resolver.CanResolveValue(toResolve) {
			return resolver
		}
	}
	return nil
}


