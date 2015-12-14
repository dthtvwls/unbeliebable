package unbeliebable

import (
	"net"
	"time"
)

type Vote struct {
	IP      net.IP
	Time    time.Time
	Against bool
}
