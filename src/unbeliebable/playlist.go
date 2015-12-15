package unbeliebable

import (
	"errors"
	"net"
	"sort"
	"time"
)

type Playlist struct {
	NowPlaying  *Song
	ElapsedTime int
	Songs       []Song
}

func (m *Playlist) Add(song Song) error {
	for i := range m.Songs {
		if m.Songs[i].ID == song.ID {
			return errors.New("song already queued")
		}
	}
	m.Songs = append(m.Songs, song)
	m.Vote(song.IP, song.ID, false)
	sort.Sort(ByVotes(m.Songs))
	return nil
}

func (m *Playlist) Vote(ip net.IP, id string, against bool) error {
	for i := range m.Songs {
		if m.Songs[i].ID == id {
			value := 1
			if against {
				value = -1
			}

			err := m.Songs[i].Vote(Vote{IP: ip, Time: time.Now(), Value: value})
			if err == nil { // don't bother to sort if the vote wasn't accepted!
				sort.Sort(ByVotes(m.Songs))
			}
			return err
		}
	}
	return errors.New("song not found")
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
