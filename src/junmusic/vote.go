package junmusic

import (
    "net"
    "time"
)

type Vote struct {
    IP net.IP
    Request Request
    Time time.Time
    Against bool
}
