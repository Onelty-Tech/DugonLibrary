package helpers

import (
	"fmt"
)

/*
Convierte un Slice de tipo interface iterando todos los indices para hacer type assertion
*/
func GetStringSlice(val any) ([]string, error) {
	var result []string

	switch list := val.(type) {
	case []string:
		return list, nil
	case []interface{}:
		for _, item := range list {
			str, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("error GetStringSlice: element is not a string: %v", item)
			}
			result = append(result, str)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("error: GetStringSlice: expected array of strings, got %T", val)
	}
}

/*
obtiene un string slice de un mapa.
*/
func GetStringSliceFromMap(event map[string]any, key string) ([]string, error) {
	var result []string
	value, exists := event[key]
	if !exists {
		return []string{}, fmt.Errorf("error GetStringSliceFromMap: missing key '%s'", key)
	}
	switch list := value.(type) {
	case []string:
		return list, nil
	case []interface{}:
		for _, item := range list {
			str, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("error GetStringSliceFromMap: element is not a string: %v", item)
			}
			result = append(result, str)
		}
		return result, nil
	default:
		return nil, fmt.Errorf("error: GetStringSliceFromMap: expected array of strings, got %T", value)
	}
}
