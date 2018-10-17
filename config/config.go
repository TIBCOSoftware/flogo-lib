// This package contains all the default values for the configuration
package config

import (
	"os"
	"strconv"
)

const (
	ENV_LOG_DATE_FORMAT_KEY       = "FLOGO_LOG_DTFORMAT"
	LOG_DATE_FORMAT_DEFAULT       = "2006-01-02 15:04:05.000"
	ENV_LOG_LEVEL_KEY             = "FLOGO_LOG_LEVEL"
	ENV_LOG_FORMAT                = "FLOGO_LOG_FORMAT"
	LOG_FORMAT_DEFAULT            = "TEXT"
	LOG_LEVEL_DEFAULT             = "INFO"
	ENV_APP_CONFIG_LOCATION_KEY   = "FLOGO_CONFIG_PATH"
	APP_CONFIG_LOCATION_DEFAULT   = "flogo.json"
	ENV_STOP_ENGINE_ON_ERROR_KEY  = "FLOGO_ENGINE_STOP_ON_ERROR"
	ENV_DATA_SECRET_KEY_KEY       = "FLOGO_DATA_SECRET_KEY"
	DATA_SECRET_KEY_DEFAULT       = "flogo"
	ENV_APP_PROPERTY_OVERRIDE_KEY = "FLOGO_APP_PROPS_OVERRIDE"
	ENV_APP_PROPERTY_RESOLVER_KEY = "FLOGO_APP_PROPS_VALUE_RESOLVER"
	ENV_LOG_LEVEL_OVERRIDE_KEY    = "FLOGO_CONTRIB_LOG_LEVEL_OVERRIDE"
)

var defaultLogLevel = LOG_LEVEL_DEFAULT

//GetFlogoConfigPath returns the flogo config path
func GetFlogoConfigPath() string {
	flogoConfigPathEnv := os.Getenv(ENV_APP_CONFIG_LOCATION_KEY)
	if len(flogoConfigPathEnv) > 0 {
		return flogoConfigPathEnv
	}
	return APP_CONFIG_LOCATION_DEFAULT
}

func SetDefaultLogLevel(logLevel string) {
	defaultLogLevel = logLevel
}

//GetLogLevel returns the log level
func GetLogLevel() string {
	logLevelEnv := os.Getenv(ENV_LOG_LEVEL_KEY)
	if len(logLevelEnv) > 0 {
		return logLevelEnv
	}
	return defaultLogLevel
}

//GetLogLevel returns the log level
func GetLogFormat() string {
	logFormatEnv := os.Getenv(ENV_LOG_FORMAT)
	if len(logFormatEnv) > 0 {
		return logFormatEnv
	}
	return LOG_FORMAT_DEFAULT
}

func GetLogDateTimeFormat() string {
	logLevelEnv := os.Getenv(ENV_LOG_DATE_FORMAT_KEY)
	if len(logLevelEnv) > 0 {
		return logLevelEnv
	}
	return LOG_DATE_FORMAT_DEFAULT
}

func StopEngineOnError() bool {
	stopEngineOnError := os.Getenv(ENV_STOP_ENGINE_ON_ERROR_KEY)
	if len(stopEngineOnError) == 0 {
		return true
	}
	b, _ := strconv.ParseBool(stopEngineOnError)
	return b
}

func GetDataSecretKey() string {
	key := os.Getenv(ENV_DATA_SECRET_KEY_KEY)
	if len(key) > 0 {
		return key
	}
	return DATA_SECRET_KEY_DEFAULT
}

func GetAppPropertiesOverride() string {
	key := os.Getenv(ENV_APP_PROPERTY_OVERRIDE_KEY)
	if len(key) > 0 {
		return key
	}
	return ""
}

func GetAppPropertiesValueResolver() string {
	key := os.Getenv(ENV_APP_PROPERTY_RESOLVER_KEY)
	if len(key) > 0 {
		return key
	}
	return ""
}

func GetContribLogLevelOverride() string {
	key := os.Getenv(ENV_LOG_LEVEL_OVERRIDE_KEY)
	if len(key) > 0 {
		return key
	}
	return ""
}
