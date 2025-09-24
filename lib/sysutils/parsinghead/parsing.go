package parsinghead

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type points struct {
	FStart bool
	Start  int
	End    int
}

const (
	END = '\n'
)

/*
Parsea el header de un paquete, mediante bufio.Reader
retorna:
[]byte: la cabecera binaria
*/
func ParseBinaryHead(reader *bufio.Reader) ([]byte, error) {
	//Lee la longitud del paquete(header)
	headerBytes, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf(`error: reading "\n": %w`, err)
	}
	//quita los saltos de linea (\n)
	headerBytes = strings.TrimSpace(headerBytes)

	//convierte a int
	cnv_headerLength, err := strconv.Atoi(headerBytes)
	if err != nil {
		return []byte{}, fmt.Errorf("error ParseBinaryHead: %w", err)
	}
	//crea el buffer
	headerBuffer := make([]byte, cnv_headerLength)
	//lee y escribe informacion en el buffer
	_, err = io.ReadFull(reader, headerBuffer)
	if err != nil {
		return []byte{}, fmt.Errorf("error ParseBinaryHead:%w", err)
	}
	fmt.Printf("\033[33mParseBinaryHead Testing Buffer:%s\033[0m\n", headerBuffer)
	return headerBuffer, nil
}

/*
Une la cabecera binaria del paquete en un solo slice de bytes
*/
func BindHeadPacket(head, container []byte) []byte {
	var buffer []byte
	buffer = append(buffer, head...)
	buffer = append(buffer, container...)
	return buffer
}
