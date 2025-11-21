package helpers

import "fmt"

/*
	De un map[string]any buscamos una key que es un map[string]any y lo pasamos a map[string]any
	todos sus elementos
*/
func GetStringMapString(m map[string]any, key string) (map[string]string, error) {
	// Inicializar el map que vamos a retornar
	newReturn := make(map[string]string)

	// Verifica si la clave existe en el map
	value, exists := m[key]
	if !exists {
		return newReturn, fmt.Errorf("error GetStringMapString: key '%s' not found", key)
	}

	// Verifica que el valor es del tipo correcto
	mapNew, ok := value.(map[string]any)
	if !ok {
		return newReturn, fmt.Errorf("error GetStringMapString: invalid data type for key '%s', expected map[string]any", key)
	}

	// Convertir los valores a string
	for k, v := range mapNew {
		strValue, ok := v.(string)
		if !ok {
			return newReturn, fmt.Errorf("error GetStringMapString: value for key '%s' is not a string", k)
		}
		newReturn[k] = strValue
	}

	return newReturn, nil
}
