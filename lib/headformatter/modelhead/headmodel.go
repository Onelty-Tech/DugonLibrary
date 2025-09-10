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
	Streaming string
}

func Unmarshal(data []byte) (Head, error) {
	DataSlice := strings.Split(strings.TrimSpace(string(data)), ",")
	if len(DataSlice) < 4 {
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
		Streaming: DataSlice[3],
	}, nil
}

func Marshal(h Head) ([]byte, error) {
	format := fmt.Sprintf("%s,%d,%s,%s", h.Type, h.Length, h.Target, h.Streaming)
	format = fmt.Sprintf("%d%c%s", len(format), '\n', format)
	return []byte(format), nil
}
