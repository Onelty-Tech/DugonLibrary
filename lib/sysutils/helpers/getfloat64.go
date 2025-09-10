package helpers

import "fmt"

/*
	Obtiene de un map[string]any tal key y la retorna
*/
func GetFloat64(m map[string]any, key string) (float64, error) {
	value, ok := m[key]
	if !ok {
		return 0, fmt.Errorf("error: missing key '%s'", key)
	}
	numf, ok := value.(float64)
	if !ok {
		return 0, fmt.Errorf("error: invalid type for key '%s': expected float64", key)
	}
	return numf, nil
}
