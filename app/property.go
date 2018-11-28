package app

import (
	"errors"
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"sync"
)

var (
	propertyProvider   *PropertyProvider
	propValueResolvers = make(map[string]PropertyValueResolver)
	lock               = &sync.Mutex{}
)


func init() {
	propertyProvider = &PropertyProvider{properties: make(map[string]interface{})}
}

func RegisterPropertyValueResolver(relType string, resolver PropertyValueResolver) error {
	lock.Lock()
	defer lock.Unlock()
	_, ok := propValueResolvers[relType]
	if ok {
		errMsg := fmt.Sprintf("Property value resolver is already registered for type - '%s'", relType)
		logger.Errorf(errMsg)
		return errors.New(errMsg)
	}
	propValueResolvers[relType] = resolver
	return nil
}

func GetPropertyValueResolver(relType string) PropertyValueResolver {
    return propValueResolvers[relType]
}

func GetPropertyProvider() *PropertyProvider {
	return propertyProvider
}

type PropertyProvider struct {
	properties map[string]interface{}
}

type PropertyValueResolver interface {
	ResolveValue(key string) (interface{}, error)
}

func (pp *PropertyProvider) GetProperty(property string) (interface{}, bool) {
	prop, exists := pp.properties[property]
	return prop, exists
}

func (pp *PropertyProvider) SetProperty(property string, value interface{}) {
	pp.properties[property] = value
}

func (pp *PropertyProvider) SetProperties(value map[string]interface{}) {
	pp.properties = value
}
