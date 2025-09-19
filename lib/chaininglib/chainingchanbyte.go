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
loop:
	for {
		select {
		case value := <-c:
			if _, err := strBuilding.Write(value); err != nil {
				return []byte{}, fmt.Errorf("error ChainingChanByte; %w", err)
			}
		default:
			break loop
		}
	}
	return []byte(strBuilding.String()), nil
}
