package environ

import (
	"log"
	"strings"
	"time"
	"unicode"

	"github.com/spf13/cast"

	"github.com/gopub/conv"
	"github.com/gopub/gox"
)

type Manager interface {
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, value interface{})
}

var DefaultManager Manager = NewViperManager()

func Has(key string) bool {
	return DefaultManager.Has(key)
}

func Get(key string) interface{} {
	return DefaultManager.Get(key)
}

func Set(key string, value interface{}) {
	DefaultManager.Set(key, value)
}

func String(key string, defaultValue string) string {
	s, err := conv.ToString(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return s
}

func MustString(key string) string {
	v, _ := conv.ToString(DefaultManager.Get(key))
	if v == "" {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Int(key string, defaultValue int) int {
	v, err := conv.ToInt(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return v
}

func MustInt(key string) int {
	v, err := conv.ToInt(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Int64(key string, defaultValue int64) int64 {
	v, err := conv.ToInt64(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return v
}

func MustInt64(key string) int64 {
	v, err := conv.ToInt64(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Float64(key string, defaultValue float64) float64 {
	v, err := conv.ToFloat64(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return v
}

func MustFloat64(key string) float64 {
	v, err := conv.ToFloat64(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Duration(key string, defaultValue time.Duration) time.Duration {
	v, err := conv.ToDuration(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return v
}

func MustDuration(key string) time.Duration {
	v, err := conv.ToDuration(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func Bool(key string, defaultValue bool) bool {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	v, err := conv.ToBool(DefaultManager.Get(key))
	if err != nil {

	}
	return v
}

func MustBool(key string) bool {
	v, err := conv.ToBool(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	return v
}

func IntSlice(key string, defaultValue []int) []int {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	v, _ := conv.ToIntSlice(DefaultManager.Get(key))
	return v
}

func MustIntSlice(key string) []int {
	v, err := conv.ToIntSlice(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}

func StringSlice(key string, defaultValue []string) []string {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	v, _ := conv.ToStringSlice(DefaultManager.Get(key))
	return v
}

func MustStringSlice(key string) []string {
	v, err := conv.ToStringSlice(DefaultManager.Get(key))
	if err != nil {
		log.Panicf("%s is not defined", key)
	}
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}

func Map(key string, defaultValue gox.M) gox.M {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	return cast.ToStringMap(DefaultManager.Get(key))
}

func MustMap(key string) gox.M {
	if !DefaultManager.Has(key) {
		log.Panicf("%s is not defined", key)
	}
	v := cast.ToStringMap(DefaultManager.Get(key))
	if len(v) == 0 {
		log.Panicf("%s is empty", key)
	}
	return v
}

func SizeInBytes(key string, defaultValue int) int {
	if !DefaultManager.Has(key) {
		return defaultValue
	}
	s, err := conv.ToString(DefaultManager.Get(key))
	if err != nil {
		return defaultValue
	}
	return int(parseSizeInBytes(s))
}

func MustSizeInBytes(key string) int {
	if !DefaultManager.Has(key) {
		log.Panicf("%s is not defined", key)
	}
	v := DefaultManager.Get(key)
	s, err := conv.ToString(v)
	if err != nil {
		log.Panicf("Cast to string %v: %v", v, err)
	}
	if len(s) == 0 {
		log.Panicf("%s is empty", key)
	}
	return int(parseSizeInBytes(s))
}

// parseSizeInBytes converts strings like 1GB or 12 mb into an unsigned integer number of bytes
func parseSizeInBytes(sizeStr string) uint {
	sizeStr = strings.TrimSpace(sizeStr)
	lastChar := len(sizeStr) - 1
	multiplier := uint(1)

	if lastChar > 0 {
		if sizeStr[lastChar] == 'b' || sizeStr[lastChar] == 'B' {
			if lastChar > 1 {
				switch unicode.ToLower(rune(sizeStr[lastChar-1])) {
				case 'k':
					multiplier = 1 << 10
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				case 'm':
					multiplier = 1 << 20
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				case 'g':
					multiplier = 1 << 30
					sizeStr = strings.TrimSpace(sizeStr[:lastChar-1])
				default:
					multiplier = 1
					sizeStr = strings.TrimSpace(sizeStr[:lastChar])
				}
			}
		}
	}

	size, _ := conv.ToInt(sizeStr)
	if size < 0 {
		size = 0
	}

	return safeMul(uint(size), multiplier)
}

func safeMul(a, b uint) uint {
	c := a * b
	if a > 1 && b > 1 && c/b != a {
		return 0
	}
	return c
}
