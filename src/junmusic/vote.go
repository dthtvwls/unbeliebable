package junmusic

import (
	"net"
	"time"
)

type Vote struct {
	IP      net.IP
	Time    time.Time
	ID      string
	Against bool
}
