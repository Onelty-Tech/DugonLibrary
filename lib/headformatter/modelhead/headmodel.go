package modelhead

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Head struct {
	//el que envio el paquete
	Sender string
	//tipo de paquete
	Type string
	//longitud a parsear
	Length int
	//objetivo del paquete(cliente)
	Target string
	//objetivo de canal
	Channel string
	//targets de streamings
	Streaming []string
}

/*
decodifica los []bytes a una estructura Head de la cabecera binaria
*/
func Unmarshal(data []byte) (Head, error) {
	var head Head
	buffer := bytes.NewBuffer(data)
	if err := gob.NewDecoder(buffer).Decode(&head); err != nil {
		return Head{}, fmt.Errorf("error Unmarshal, library:'DugonLibrary/modelhead': %w", err)
	}
	return head, nil
}

/*
Convierte una cabecera binaria a []bytes
*/
func Marshal(h Head) ([]byte, error) {
	//creamos un buffer
	var buffer bytes.Buffer

	//señalamos la codificacion a la variable "buffer" y que tome los datos de la struct "Head"
	if err := gob.NewEncoder(&buffer).Encode(&h); err != nil {
		return []byte{}, fmt.Errorf("error Marshal in modelhead library:%w", err)
	}
	format := fmt.Sprintf("%d%c%s", len(buffer.Bytes()), '\n', buffer.Bytes())
	//devolvemos solo el array de bytes
	return []byte(format), nil
}
