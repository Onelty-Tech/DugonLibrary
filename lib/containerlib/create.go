package containerlib

import (
	"ship/api/src/dock/modeldock"
	"time"

	"github.com/Onelty-Tech/DugonLibrary/lib/sysutils/typenames"

	"github.com/Onelty-Tech/DugonModels/modelcontainer"
)

// Crea un Container
func CreateContainer(
	dock *modeldock.Dock,
	packagename string,
	target string,
	channel string, //igualmente posiblemente tenga su propia funcion para agregarlo al contenedor mas adelante
	content string,
	typeupdate bool,
) modelcontainer.Container {
	return modelcontainer.Container{
		Header: modelcontainer.Header{
			HeaderName:    packagename,
			HeaderSender:  dock.DockName,
			HeaderTarget:  target,
			HeaderChannel: channel,
			HeaderDate:    time.Now().Format(time.RFC3339),
			Pack: modelcontainer.Pack{
				Package: typenames.PackagesType[1],
			},
		},
		Options: modelcontainer.Options{
			OptionsTypeUpdate: typeupdate,
		},
		Body: modelcontainer.Body{BodyContent: content},
	}
}
