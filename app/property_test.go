package app

import (
	"testing"
	"os"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvValueResolver(t *testing.T) {
	os.Setenv(config.ENV_APP_PROPERTY_RESOLVER_KEY, "env")
	os.Setenv("TEST_PROP", "testprop")
	defer func() {
		os.Unsetenv(config.ENV_APP_PROPERTY_RESOLVER_KEY)
		os.Unsetenv("TEST_PROP")
	}()

	resolver := GetPropertyValueResolver("env")
	assert.NotNil(t, resolver)
	resolvedVal, found := resolver.LookupValue("TEST_PROP")
	assert.True(t, true, found)
	assert.Equal(t, "testprop", resolvedVal)

	_, found = resolver.LookupValue("TEST_PROP1")
	assert.False(t, false, found)
}