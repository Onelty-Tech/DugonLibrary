package containerlib

import (
	"encoding/json"
	"fmt"

	"github.com/Onelty-Tech/DugonModels/modelcontainer"
)

/*
Genera el contenedor mediante []byte
*/
func GenerateByBuffer(buffer *[]byte) (modelcontainer.Container, error) {
	var Container modelcontainer.Container
	err := json.Unmarshal(*buffer, &Container)
	if err != nil {
		return modelcontainer.Container{}, fmt.Errorf("error: Could not deserialize container data;%s", err)
	}
	return Container, nil
}
