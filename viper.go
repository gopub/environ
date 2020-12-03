package environ

import (
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

type ViperManager struct {
	dotKeyPattern        *regexp.Regexp
	underscoreKeyPattern *regexp.Regexp
}

func NewViperManager() *ViperManager {
	viper.AutomaticEnv()
	m := &ViperManager{}
	m.dotKeyPattern = regexp.MustCompile("[a-zA-Z][a-zA-Z0-9]*(\\.[a-zA-Z][a-zA-Z0-9]*)+")
	m.underscoreKeyPattern = regexp.MustCompile("[a-zA-Z][a-zA-Z0-9]*(_[a-zA-Z][a-zA-Z0-9]*)+")
	return m
}

func (m *ViperManager) Has(key string) bool {
	return viper.IsSet(key)
}

func (m *ViperManager) Get(key string) interface{} {
	if v := viper.Get(key); v != nil {
		return v
	}
	if m.dotKeyPattern.MatchString(key) {
		key = strings.Replace(key, ".", "_", -1)
		return viper.Get(key)
	}

	if m.underscoreKeyPattern.MatchString(key) {
		key = strings.Replace(key, "_", ".", -1)
		return viper.Get(key)
	}

	return nil
}

func (m *ViperManager) Set(key string, value interface{}) {
	viper.Set(key, value)
}
