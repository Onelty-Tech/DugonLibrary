package chaininglib

import (
	"fmt"
	"strings"
)

/*
encadena todos los paquetes actuales de un chan string en un string.
*/
func ChainingChanByte(c chan []byte) ([]byte, error) {
	var strBuilding strings.Builder
	for b := range c {
		if _, err := strBuilding.Write(b); err != nil {
			return []byte{}, fmt.Errorf("error ChainingChanString: Could not write byte in strings.Builder;%w", err)
		}
	}
	return []byte(strBuilding.String()), nil
}
