package netpipe

import (
	"log"
	"net"

	"github.com/aln730/tun-based-vpn/internal/crypto"
)

type Pipe struct {
	conn *net.UDPConn
	box  *crypto.Box
	peer *net.UDPAddr
}

// Creates a UDP socket bound to localAddr and prepares to talk to remoteAddr
func NewPipe(localAddr, remoteAddr string, box *crypto.Box) (*Pipe, error) {
	laddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		return nil, err
	}

	raddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return nil, err
	}

	return &Pipe{
		conn: conn,
		box:  box,
		peer: raddr,
	}, nil
}

// Send encrypted payload to peer
func (p *Pipe) Send(data []byte) error {
	encrypted, err := p.box.Encrypt(data)
	if err != nil {
		return err
	}
	_, err = p.conn.WriteToUDP(encrypted, p.peer)
	return err
}

// Returns the plaintext of a UDP packet
func (p *Pipe) Receive(buf []byte) ([]byte, int, error) {
	n, _, err := p.conn.ReadFromUDP(buf)
	if err != nil {
		return nil, 0, err
	}

	plaintext, err := p.box.Decrypt(buf[:n])
	if err != nil {
		log.Printf("decrypt error: %v", err)
		return nil, 0, err
	}

	return plaintext, len(plaintext), nil
}
