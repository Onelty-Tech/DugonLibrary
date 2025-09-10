package helpers

import "fmt"

/*
	Obtiene de un map[string]any tal key y la retorna
*/
func GetString(m map[string]any, key string) (string, error) {
	value, ok := m[key]
	if !ok {
		return "", fmt.Errorf("error: missing key '%s'", key)
	}
	data, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("error: invalid type for key '%s': expected string", key)
	}
	return data, nil
}
