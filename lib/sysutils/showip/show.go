package showip

import (
	"fmt"
	"net"
)

/*
-----------------------------
Libreria para obtener la IPv4
-----------------------------
*/
func GetIPv4() string {
	//Obtiene todas las Direcciones IP
	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("error getting ip addresses:", err)
		return ""
	}

	//Itera sobre todas las IP
	for _, addr := range address {
		//ase un Type Assertion
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
