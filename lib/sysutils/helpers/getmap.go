package helpers

import "fmt"

/*
	Obtiene de una key un map[string]any
*/
func GetMapAny(m map[string]any, key string) (map[string]any, error) {
	val, ok := m[key]
	if !ok {
		return nil, fmt.Errorf("error GetMapAny: missing key '%s'", key)
	}
	valCnv, ok := val.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("error GetMapAny: Invalid data type for key '%s': expected map[string]any", key)
	}
	return valCnv, nil
}
