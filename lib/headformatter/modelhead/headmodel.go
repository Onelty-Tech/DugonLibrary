package modelhead

import (
	"fmt"
	"strconv"
	"strings"
)

type Head struct {
	Type      string
	Length    int
	Target    string
	Channel   string
	Streaming string
}

/*
decodifica los []bytes a una estructura Head de la cabecera binaria
*/
func Unmarshal(data []byte) (Head, error) {
	DataSlice := strings.Split(strings.TrimSpace(string(data)), ",")
	if len(DataSlice) < 5 {
		return Head{}, fmt.Errorf("error: Invalid length, missing arguments in Unmarshal headmodel;")
	}
	Length, err := strconv.Atoi(DataSlice[1])
	if err != nil {
		return Head{}, fmt.Errorf("error: Could not unmarshal;%s", err)
	}
	return Head{
		Type:      DataSlice[0],
		Length:    Length,
		Target:    DataSlice[2],
		Channel:   DataSlice[3],
		Streaming: DataSlice[4],
	}, nil
}

/*
Convierte una cabecera binaria a []bytes
*/
func Marshal(h Head) ([]byte, error) {
	//agrega la cabecera binaria(struct) a un string separado por comas para facilitar la decodificacion de este
	format := fmt.Sprintf("%s,%d,%s,%s,%s", h.Type, h.Length, h.Target, h.Channel, h.Streaming)
	//agrega primero la longitud para luego al decodificarlo se pueda obtener directamente por buffer
	format = fmt.Sprintf("%d%c%s", len(format), '\n', format)
	return []byte(format), nil
}
