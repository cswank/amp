package amp

import (
	"fmt"
	"net"
)

const (
	packetTpl = `<?xml version="1.0" encoding="utf-8"?>
<%s>
  %%s
</%s>`
)

type Amp struct {
	ctrl *connection
}

func New() (Amp, error) {
	ctrl, err := newConnection(7002, "emotivaControl")
	return Amp{ctrl: ctrl}, err
}

func (a Amp) Close() {
	a.ctrl.close()
}

func (a Amp) On() error {
	return a.ctrl.write(`<power_on value="0" ack="no" />`)
}

func (a Amp) Off() error {
	return a.ctrl.write(`<power_off value="0" ack="no" />`)
}

type connection struct {
	addr   *net.UDPAddr
	conn   *net.UDPConn
	packet string
}

func (c connection) write(msg string) error {
	_, err := c.conn.WriteToUDP([]byte(fmt.Sprintf(c.packet, msg)), c.addr)
	return err
}

func (c connection) close() {
	c.conn.Close()
}

func newConnection(port int, packet string) (*connection, error) {
	a, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("255.255.255.255:%d", port))
	if err != nil {
		return nil, err
	}

	c, err := net.ListenUDP("udp4", a)
	if err != nil {
		return nil, err
	}

	return &connection{addr: a, conn: c, packet: fmt.Sprintf(packetTpl, packet, packet)}, nil
}
