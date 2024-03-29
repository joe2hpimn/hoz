package hoz

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

type socks5 interface {
	handshakeSocks([]byte) (bool, []byte, error)
	parseSocks([]byte) (bool, []byte, error)
}

func parseSocks5Request(b []byte) ([]byte, bool) {
	n := len(b)
	resp := b[:0]
	ver := b[0]
	cmd := b[1]
	rsv := b[2]
	atyp := b[3]
	// only support tcp
	resp = append(resp, ver)
	// success
	resp = append(resp, 0x00)
	/*X'00' succeeded
	X'01' general SOCKS server failure
	X'02' connection not allowed by ruleset
	X'03' Network unreachable
	X'04' Host unreachable
	X'05' Connection refused
	X'06' TTL expired
	X'07' Command not supported
	X'08' Address type not supported
	X'09' to X'FF' unassigned*/
	resp = append(resp, rsv)
	resp = append(resp, atyp)
	if cmd == 1 {
		var host, port string
		switch b[3] {
		case 0x01: //IP V4
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03: //Domain
			host = string(b[5 : n-2]) //b[4] domain length
		case 0x04: //IP V6
			host = net.IP{b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15], b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))
		//LOG.Printf("type %d， target host %s port %s\n", atyp, string(host), port)
		// socks to http, send to remote
		to5 := to5Connect(host, port)
		return to5, true
	} else {
		// failed
		resp[1] = 0x01
		resp = append(resp, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}...)
	}
	return resp, false
}

func to5Connect(host, port string) []byte {
	var bf bytes.Buffer
	bf.WriteString("SOCKS5 /socks5 HTTP/1.1\r\n")
	bf.WriteString(fmt.Sprintf("Shost: %s\r\n", host))
	bf.WriteString(fmt.Sprintf("Sport: %s\r\n", port))
	bf.WriteString("\r\n")
	return bf.Bytes()
}
