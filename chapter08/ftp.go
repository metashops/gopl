package main

import (
	`fmt`
	`net`
	`strconv`
)

type conn struct {
	rw           net.Conn // "Protocol Interpreter" connection
	dataHostPort string
	prevCmd      string
	pasvListener net.Listener
	cmdErr       error // Saved command connection write error.
	binary       bool
}

func NewConn(cmdConn net.Conn) *conn {
	return &conn{rw: cmdConn}
}

// hostPortToFTP returns a comma-separated, FTP-style address suitable for
// replying to the PASV command.
func hostPortToFTP(hostport string) (addr string, err error) {
	host, portStr, err := net.SplitHostPort(hostport)
	if err != nil {
		return "", nil
	}
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "", err
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return "", err
	}
	ip := ipAddr.IP.To4()
	s := fmt.Sprintf("%d,%d,%d,%d,%d,%d", ip[0], ip[1], ip[2], ip[3], port/256, port%256)
	return s, nil
}

func hostPortFromFTP(address string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Sscanf(address, "%d,%d,%d,%d,%d,%d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d:%d", a, b, c, d, 256*p1+p2), nil
}
func main() {

}
