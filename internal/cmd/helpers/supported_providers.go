package helpers

import (
	"fmt"
	"github.com/davidalpert/go-yoss/internal/provider"
	"github.com/davidalpert/go-yoss/internal/provider/paramstore"
	"sort"
	"strings"
)

var SupportedProviders map[string]provider.NewProviderFn

func ValidateProviderKey(key string) error {
	if _, ok := SupportedProviders[key]; !ok {
		return fmt.Errorf("unrecognized provider %#v: supported provider are: %#v", key, strings.Join(supportedProviderKeys(), ", "))
	}
	return nil
}

func supportedProviderKeys() []string {
	result := make([]string, len(SupportedProviders))
	i := 0
	for key, _ := range SupportedProviders {
		result[i] = key
		i++
	}
	sort.Strings(result)
	return result

}
func init() {
	SupportedProviders = map[string]provider.NewProviderFn{
		paramstore.ProviderKey: paramstore.NewProvider,
	}
}
