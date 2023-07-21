package device

import (
	"net"
	"strings"
)

func GetBoundIp() string {
	conn, _ := net.Dial("udp", "8.8.8.8:53")
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}
