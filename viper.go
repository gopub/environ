package environ

import (
	"github.com/spf13/viper"
)

type ViperManager struct {
}

func NewViperManager() *ViperManager {
	viper.AutomaticEnv()
	return &ViperManager{}
}

func (m *ViperManager) Has(key string) bool {
	return viper.IsSet(key)
}

func (m *ViperManager) Get(key string) interface{} {
	return viper.Get(key)
}

func (m *ViperManager) Set(key string, value interface{}) {
	viper.Set(key, value)
}
