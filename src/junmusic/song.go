package junmusic

import (
	"net"
	"time"
)

type Song struct {
	IP               net.IP
	Time             time.Time
	ID, Name, Artist string
}
