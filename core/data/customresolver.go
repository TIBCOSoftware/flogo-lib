package data

import "sync"

var (
 resolvers []CustomResolver
 lock = &sync.Mutex{}
)


type CustomResolver interface {
	CanResolve(toResolve string) bool
	Resolve(toResolve string, scope Scope) (interface{}, error)
}


func RegisterCustomResolver(resolver CustomResolver) {
	lock.Lock()
	defer lock.Unlock()
	resolvers = append(resolvers, resolver)
}


func GetCustomResolver(toResolve string) CustomResolver  {
	for _, resolver := range resolvers {
		if resolver.CanResolve(toResolve) {
			return resolver
		}
	}
	return nil
}


