package chaininglib

import (
	"fmt"
	"strings"
)

/*
encadena todos los paquetes actuales de un chan string en un string
*/
func ChainingChanString(c chan string) (string, error) {
	var strBuilding strings.Builder
	for str := range c {
		if _, err := strBuilding.WriteString(str); err != nil {
			return "", fmt.Errorf("error ChainingChanString: Could not write string in builder string;%w", err)
		}
	}
	return strBuilding.String(), nil
}
