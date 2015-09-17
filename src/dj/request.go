package dj

import "net"

type Request struct {
    IP net.IP
    Song Song
}
