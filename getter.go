package mconfig

import (
	"strings"
	"time"
)

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} { return v.Get(key) }
func (v *Viper) Get(key string) interface{} {
	val := v.find(key)
	if val == nil {
		return nil
	}

	return val
}

// GetString returns the value associated with the key as a string.
func GetString(key string) (string, error) { return v.GetString(key) }
func (v *Viper) GetString(key string) (string, error) {
	value := v.Get(key)
	if value == nil {
		return "", ErrKeyNotExist
	}
	return ToStringE(value)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) (bool, error) { return v.GetBool(key) }
func (v *Viper) GetBool(key string) (bool, error) {
	value := v.Get(key)
	if value == nil {
		return false, ErrKeyNotExist
	}
	return ToBoolE(value)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) (int, error) { return v.GetInt(key) }
func (v *Viper) GetInt(key string) (int, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}

	return ToIntE(value)
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) (int32, error) { return v.GetInt32(key) }
func (v *Viper) GetInt32(key string) (int32, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToInt32E(v.Get(key))
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) (int64, error) { return v.GetInt64(key) }
func (v *Viper) GetInt64(key string) (int64, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToInt64E(v.Get(key))
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) (uint, error) { return v.GetUint(key) }
func (v *Viper) GetUint(key string) (uint, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}

	return ToUintE(value)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) (uint32, error) { return v.GetUint32(key) }
func (v *Viper) GetUint32(key string) (uint32, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToUint32E(v.Get(key))
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) (uint64, error) { return v.GetUint64(key) }
func (v *Viper) GetUint64(key string) (uint64, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToUint64E(v.Get(key))
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) (float64, error) { return v.GetFloat64(key) }
func (v *Viper) GetFloat64(key string) (float64, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToFloat64E(v.Get(key))
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) (time.Time, error) { return v.GetTime(key) }
func (v *Viper) GetTime(key string) (time.Time, error) {
	value := v.Get(key)
	if value == nil {
		return time.Time{}, ErrKeyNotExist
	}
	return ToTimeE(v.Get(key))
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) (time.Duration, error) { return v.GetDuration(key) }
func (v *Viper) GetDuration(key string) (time.Duration, error) {
	value := v.Get(key)
	if value == nil {
		return 0, ErrKeyNotExist
	}
	return ToDurationE(v.Get(key))
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) ([]int, error) { return v.GetIntSlice(key) }
func (v *Viper) GetIntSlice(key string) ([]int, error) {
	value := v.Get(key)
	if value == nil {
		return nil, ErrKeyNotExist
	}
	return ToIntSliceE(v.Get(key))
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) ([]string, error) { return v.GetStringSlice(key) }
func (v *Viper) GetStringSlice(key string) ([]string, error) {
	value := v.Get(key)
	if value == nil {
		return nil, ErrKeyNotExist
	}
	return ToStringSliceE(v.Get(key))
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) (map[string]interface{}, error) { return v.GetStringMap(key) }
func (v *Viper) GetStringMap(key string) (map[string]interface{}, error) {
	value := v.Get(key)
	if value == nil {
		return nil, ErrKeyNotExist
	}
	return ToStringMapE(v.Get(key))
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) (map[string]string, error) { return v.GetStringMapString(key) }
func (v *Viper) GetStringMapString(key string) (map[string]string, error) {
	value := v.Get(key)
	if value == nil {
		return nil, ErrKeyNotExist
	}
	return ToStringMapStringE(v.Get(key))
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) (map[string][]string, error) {
	return v.GetStringMapStringSlice(key)
}
func (v *Viper) GetStringMapStringSlice(key string) (map[string][]string, error) {
	value := v.Get(key)
	if value == nil {
		return nil, ErrKeyNotExist
	}
	return ToStringMapStringSliceE(v.Get(key))
}

// Given a key, find the value.
//
// Viper will check to see if an alias exists first.
// Viper will then check in the following order:
// flag, env, config file, key/value store.
// Lastly, if no value was found and flagDefault is true, and if the key
// corresponds to a flag, the flag's default value is returned.
//
// Note: this assumes a lower-cased key given.
func (v *Viper) find(lcaseKey string) interface{} {
	var (
		val  interface{}
		path = strings.Split(lcaseKey, v.keyDelim)
	)

	// Config file next
	val = v.searchMapWithPathPrefixes(v.config, path)
	if val != nil {
		return val
	}

	return nil
}

// searchMapWithPathPrefixes recursively searches for a value for path in source map.
//
// While searchMap() considers each path element as a single map key, this
// function searches for, and prioritizes, merged path elements.
// e.g., if in the source, "foo" is defined with a sub-key "bar", and "foo.bar"
// is also defined, this latter value is returned for path ["foo", "bar"].
//
// This should be useful only at config level (other maps may not contain dots
// in their keys).
//
// Note: This assumes that the path entries and map keys are lower cased.
func (v *Viper) searchMapWithPathPrefixes(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}

	// search for path prefixes, starting from the longest one
	for i := len(path); i > 0; i-- {
		prefixKey := strings.Join(path[0:i], v.keyDelim)

		next, ok := source[prefixKey]
		if ok {
			// Fast path
			if i == len(path) {
				return next
			}

			// Nested case
			var val interface{}
			switch next.(type) {
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				val = v.searchMapWithPathPrefixes(next.(map[string]interface{}), path[i:])
			default:
				// got a value but nested key expected, do nothing and look for next prefix
			}
			if val != nil {
				return val
			}
		}
	}

	// not found
	return nil
}
