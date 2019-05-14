package configHelper

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/pelletier/go-toml"
)

func GetTree(tree *toml.Tree, key string) *toml.Tree {
	if tree.Has(key) {
		if value, ok := tree.Get(key).(*toml.Tree); ok {
			return value
		}
	}
	return new(toml.Tree)
}

func GetString(tree *toml.Tree, key string, values ...string) string {
	if tree.Has(key) {
		return fmt.Sprintf("%v", tree.Get(key))
	} else if len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetFloat64(tree *toml.Tree, key string, values ...float64) float64 {
	if tree.Has(key) {
		value := tree.Get(key)
		switch value.(type) {
		case float64:
			return value.(float64)
		case int64:
			return float64(value.(int64))
		case uint64:
			return float64(value.(uint64))
		case string:
			if value, err := strconv.ParseFloat(value.(string), 64); err == nil {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0.0
}

func GetInt64(tree *toml.Tree, key string, values ...int64) int64 {
	if tree.Has(key) {
		value := tree.Get(key)
		switch value.(type) {
		case int64:
			return value.(int64)
		case uint64:
			return int64(value.(uint64))
		case float64:
			return int64(value.(float64))
		case string:
			if value, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetUint64(tree *toml.Tree, key string, values ...uint64) uint64 {
	if tree.Has(key) {
		value := tree.Get(key)
		switch value.(type) {
		case uint64:
			return value.(uint64)
		case int64:
			return uint64(value.(int64))
		case float64:
			return uint64(value.(float64))
		case string:
			if value, err := strconv.ParseUint(value.(string), 10, 64); err == nil {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func GetBool(tree *toml.Tree, key string, values ...bool) bool {
	if tree.Has(key) {
		value := tree.Get(key)
		switch value.(type) {
		case bool:
			return value.(bool)
		case string:
			if value, err := strconv.ParseBool(value.(string)); err == nil {
				return value
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return false
}

func GetDuration(tree *toml.Tree, key string, values ...time.Duration) time.Duration {
	if tree.Has(key) {
		value := tree.Get(key)
		switch value.(type) {
		case string:
			duration, err := time.ParseDuration(value.(string))
			if err != nil {
				log.Println(err)
			} else {
				return duration
			}
		}
	}
	if len(values) > 0 {
		return values[0]
	}
	return 0 * time.Second
}

func GetStringArray(tree *toml.Tree, key string, values ...[]string) []string {
	strings := make([]string, 0)
	if tree.Has(key) {
		if array, ok := tree.Get(key).([]interface{}); ok {
			for _, value := range array {
				strings = append(strings, fmt.Sprintf("%v", value))
			}
		}
	}
	if len(strings) == 0 && len(values) > 0 {
		return values[0]
	}
	return strings
}
