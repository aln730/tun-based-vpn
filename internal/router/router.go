package router

import (
	"io"
	"log"

	"github.com/aln730/tun-based-vpn/internal/netpipe"
	"github.com/songgao/water"
)

type Router struct {
	ifce *water.Interface
	pipe *netpipe.Pipe
}

func New(ifce *water.Interface, pipe *netpipe.Pipe) *Router {
	return &Router{
		ifce: ifce,
		pipe: pipe,
	}
}

func (r *Router) Start() {
	go r.tunToPipe()
	go r.pipeToTun()
}

func (r *Router) tunToPipe() {
	buf := make([]byte, 2000)
	for {
		n, err := r.ifce.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("[tunToPipe] read from tun error: %v", err)
			continue
		}

		err = r.pipe.Send(buf[:n])
		if err != nil {
			log.Printf("[tunToPipe] send to pipe error: %v", err)
		} else {
			log.Printf("[tunToPipe] sent %d bytes", n)
		}
	}
}

func (r *Router) pipeToTun() {
	buf := make([]byte, 2000)
	for {
		data, n, err := r.pipe.Receive(buf)
		if err != nil {
			log.Printf("[pipeToTun] receive from pipe error: %v", err)
			continue
		}

		_, err = r.ifce.Write(data[:n])
		if err != nil {
			log.Printf("[pipeToTun] write to tun error: %v", err)
		} else {
			log.Printf("[pipeToTun] wrote %d bytes to tun", n)
		}
	}
}
