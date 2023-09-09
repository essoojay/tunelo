package wire

import (
	"fmt"
	"net"
	"time"

	"tunelo/transport"
)

func (w *Wire) readVPNResponse(vpnConn net.Conn, transportConn *transport.Conn) {
	buf := make([]byte, w.BufSize)

	for {
		readDeadline := time.Now().Add(10 * time.Second)
		vpnConn.SetReadDeadline(readDeadline)

		n, err := vpnConn.Read(buf)
		if err != nil {
			w.Logger.Error(fmt.Errorf("reading from vpn conn: %v", err), nil)
			break
		}

		w.Logger.Info("read vpn response.", nil)

		go w.UDPMsgHandler(transportConn, buf[:n])
	}

	vpnConn.Close()

	w.Logger.Info("closed vpn udp conn.", nil)
}