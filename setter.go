package mconfig

import (
	"fmt"
	"strings"
)

// copyAndInsensitiviseMap behaves like insensitiviseMap, but creates a copy of
// any map it makes case insensitive.
func copyAndInsensitiviseMap(m map[string]interface{}) map[string]interface{} {
	nm := make(map[string]interface{})

	for key, val := range m {
		switch v := val.(type) {
		case map[string]interface{}:
			nm[key] = copyAndInsensitiviseMap(v)
		default:
			nm[key] = v
		}
	}

	return nm
}

// toCaseInsensitiveValue checks if the value is a  map;
// if so, create a copy and lower-case the keys recursively.
func toCaseInsensitiveValue(value interface{}) interface{} {
	switch v := value.(type) {
	case map[string]interface{}:
		value = copyAndInsensitiviseMap(v)
	}

	return value
}

// Set sets the value for the key in the override register.
// Set is case-insensitive for a key.
// Will be used instead of values obtained via
// flags, config file, ENV, default, or key/value store.
func Set(key string, value interface{}) error { return v.Set(key, value) }
func (v *Viper) Set(key string, value interface{}) error {
	// If alias passed in, then set the proper override
	value = toCaseInsensitiveValue(value)

	path := strings.Split(key, v.keyDelim)
	err := v.setMapWithPathPrefixes(v.config, path, value)
	if err != nil {
		return fmt.Errorf("set value failed:%v", err)
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
func (v *Viper) setMapWithPathPrefixes(source map[string]interface{}, path []string, value interface{}) error {
	if len(path) == 0 {
		return fmt.Errorf("invalid arg")
	} else if len(path) == 1 {
		// final point, no matter exist or not, replace it imediately

		old, ok := source[path[0]]
		if ok { // already exist, if the same, do nothing, else replace it and set isChanged flag
			needRepalce := false
			switch value.(type) {
			case float64:
				v, err := ToFloat64E(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case float32:
				v, err := ToFloat64E(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case int:
				v, err := ToIntE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case int64:
				v, err := ToIntE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case int32:
				v, err := ToIntE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case int16:
				v, err := ToIntE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case int8:
				v, err := ToIntE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case uint:
				v, err := ToUintE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case uint64:
				v, err := ToUintE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case uint32:
				v, err := ToUintE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case uint16:
				v, err := ToUintE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case uint8:
				v, err := ToUintE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case string:
				v, err := ToStringE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			case bool:
				v, err := ToBoolE(old)
				if err != nil || v != value {
					needRepalce = true
				}
				break
			default:
				needRepalce = true
			}
			if needRepalce {
				source[path[0]] = value
				v.isChanged = true
			}
		} else {
			source[path[0]] = value
		}
	} else {
		firstKey := path[0]
		next, ok := source[firstKey]
		if ok {
			// Nested case
			switch next.(type) {
			case map[string]interface{}:
				// Type assertion is safe here since it is only reached
				// if the type of `next` is the same as the type being asserted
				return v.setMapWithPathPrefixes(next.(map[string]interface{}), path[1:], value)
			default:
				// got a value but nested key expected, do nothing and look for next prefix
				return fmt.Errorf("invalid json format")
			}
		} else { // not exist
			// Nested case: first create node
			node := make(map[string]interface{})
			source[firstKey] = node
			return v.setMapWithPathPrefixes(node, path[1:], value)
		}
	}

	return nil
}
