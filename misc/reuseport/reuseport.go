package reuseport

import (
	"github.com/libp2p/go-reuseport"
)

var (
	Available    = reuseport.Available
	Control      = reuseport.Control
	Dial         = reuseport.Dial
	DialTimeout  = reuseport.DialTimeout
	Listen       = reuseport.Listen
	ListenPacket = reuseport.ListenPacket
	ResolveAddr  = reuseport.ResolveAddr
)
