package unbeliebable

import (
	"errors"
	"net"
)

type Playlist struct {
	NowPlaying  *Song
	ElapsedTime int
	Songs       []Song
}

func (m *Playlist) Add(song Song) {
	m.Songs = append(m.Songs, song)
}

func (m *Playlist) Vote(ip net.IP, id string) {

}

func (m *Playlist) Shift() (Song, error) {
	if len(m.Songs) == 0 {
		return Song{}, errors.New("empty list")
	}
	m.NowPlaying = &m.Songs[0]
	list := m.Songs[1:]
	m.Songs = make([]Song, len(list))
	copy(m.Songs, list)
	return *m.NowPlaying, nil
}
