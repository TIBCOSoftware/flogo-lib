package app

import (
	"encoding/json"

	"os"

	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"



	"github.com/TIBCOSoftware/flogo-lib/app/resource"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/TIBCOSoftware/flogo-lib/util"

)

// Config is the configuration for the App
type Config struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	Description string `json:"description"`

	Properties []*data.Attribute  `json:"properties"`
	Channels   []string           `json:"channels"`
	Triggers   []*trigger.Config  `json:"triggers"`
	Resources  []*resource.Config `json:"resources"`
	Actions    []*action.Config   `json:"actions"`
}

var appName, appVersion string

// Returns name of the application
func GetName() string {
	return appName
}

// Returns version of the application
func GetVersion() string {
	return appVersion
}

// defaultConfigProvider implementation of ConfigProvider
type defaultConfigProvider struct {
}

// ConfigProvider interface to implement to provide the app configuration
type ConfigProvider interface {
	GetApp() (*Config, error)
}

// DefaultSerializer returns the default App Serializer
func DefaultConfigProvider() ConfigProvider {
	return &defaultConfigProvider{}
}

// GetApp returns the app configuration
func (d *defaultConfigProvider) GetApp() (*Config, error) {
	return LoadConfig("")
}

func LoadConfig(flogoJson string) (*Config, error) {
	app := &Config{}
	if flogoJson == "" {
		configPath := config.GetFlogoConfigPath()

		flogo, err := os.Open(configPath)
		if err != nil {
			return nil, err
		}

		file, err := ioutil.ReadAll(flogo)
		if err != nil {
			return nil, err
		}

		updated, err := preprocessConfig(file)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(updated, &app)
		if err != nil {
			return nil, err
		}

	} else {
		updated, err := preprocessConfig([]byte(flogoJson))
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(updated, &app)
		if err != nil {
			return nil, err
		}
	}
	appName = app.Name
	appVersion = app.Version
	return app, nil
}

func preprocessConfig(appJson []byte) ([]byte, error) {

	// For now decode secret values
	re := regexp.MustCompile("SECRET:[^\\\\\"]*")
	for _, match := range re.FindAll(appJson, -1) {
		decodedValue, err := resolveSecretValue(string(match))
		if err != nil {
			return nil, err
		}
		appstring := strings.Replace(string(appJson), string(match), decodedValue, -1)
		appJson = []byte(appstring)
	}

	return appJson, nil
}

func resolveSecretValue(encrypted string) (string, error) {
	encodedValue := string(encrypted[7:])
	decodedValue, err := data.GetSecretValueHandler().DecodeValue(encodedValue)
	if err != nil {
		return "", err
	}
	return decodedValue, nil
}

func GetProperties(properties []*data.Attribute) (map[string]interface{}, error) {

	props := make(map[string]interface{})
	if properties != nil {
		overriddenProps, err := loadExternalProperties(properties)
		if err != nil {
			return props, err
		}
		for _, property := range properties {
			pValue := property.Value()
			if newValue, ok := overriddenProps[property.Name()]; ok {
				pValue = newValue
			}
			value, err := data.CoerceToValue(pValue, property.Type())
			if err != nil {
				return props, err
			}

			strVal, ok := value.(string)
			if ok && strings.HasPrefix(strVal, "SECRET:") {
				// Resolve secret value
				newVal, err := resolveSecretValue(value.(string))
				if err != nil {
					return nil, err
				}
				props[property.Name()] = newVal
			} else {
				props[property.Name()] = value
			}
		}
		return props, nil
	}

	return props, nil
}

func loadExternalProperties(properties []*data.Attribute) (map[string]interface{}, error) {

	props := make(map[string]interface{})
	propFile := config.GetAppPropertiesOverride()
	if propFile != "" {
		logger.Infof("'%s' is set. Loading overridden properties", config.ENV_APP_PROPERTY_OVERRIDE_KEY)
		//TODO move to file resolver
		if strings.HasSuffix(propFile, ".json") {
			// Override through file
			file, e := ioutil.ReadFile(propFile)
			if e != nil {
				return nil, e
			}
			e = json.Unmarshal(file, &props)
			if e != nil {
				return nil, e
			}
		} else if strings.ContainsRune(propFile, '=') {
			// Override through P1=V1,P2=V2
			overrideProps := util.ParseKeyValuePairs(propFile)
			for k, v := range overrideProps {
				props[k] = v
			}
		}
	}

	resolverType := config.GetAppPropertiesValueResolver()
	if resolverType != "" {
		logger.Infof("'%s' is set to '%s'. ", config.ENV_APP_PROPERTY_RESOLVER_KEY, resolverType)
		resolver := GetPropertyValueResolver(resolverType)
		if resolver == nil {
			errMag := fmt.Sprintf("Unsupported resolver type - %s. Resolver not registered.", resolverType)
			return nil, errors.New(errMag)
		}

		if len(props) > 0 {
			// Get value using overridden property name
			for k, v := range props {
				strVal, ok := v.(string)
				if ok {
					if len(strVal) > 0 && strVal[0] == '$' {
						// Use resolver
						newVal, found := resolver.LookupValue(strVal[1:])
						if found {
							props[k] = newVal
						} else {
							logger.Warnf("Property '%s' could not be resolved using resolver '%s'. Using default value.", strVal[1:], resolverType)
						}
					}
				} else {
					props[k] = v
				}
			}
		} else {
			// Resolver is set. Get values using app prop name
			for _, prop := range properties {
				newVal, found := resolver.LookupValue(prop.Name())
				if found {
					// Use new value
					props[prop.Name()] = newVal
				} else {
					logger.Warnf("Property - '%s' could not be resolved using resolver - '%s'. Using default value.", prop.Name(), resolverType)
				}
			}
		}
	}

	return props, nil
}

//used for old action config

//func FixUpApp(cfg *Config) {
//
//	if cfg.Resources != nil || cfg.Actions == nil {
//		//already new app format
//		return
//	}
//
//	idToAction := make(map[string]*action.Config)
//	for _, act := range cfg.Actions {
//		idToAction[act.Id] = act
//	}
//
//	for _, trg := range cfg.Triggers {
//		for _, handler := range trg.Handlers {
//
//			oldAction := idToAction[handler.ActionId]
//
//			newAction := &action.Config{Ref: oldAction.Ref}
//
//			if oldAction != nil {
//				newAction.Mappings = oldAction.Mappings
//			} else {
//				if handler.ActionInputMappings != nil {
//					newAction.Mappings = &data.IOMappings{}
//					newAction.Mappings.Input = handler.ActionInputMappings
//					newAction.Mappings.Output = handler.ActionOutputMappings
//				}
//			}
//
//			newAction.Data = oldAction.Data
//			newAction.Metadata = oldAction.Metadata
//
//			handler.Action = newAction
//		}
//	}
//
//	cfg.Actions = nil
//}
