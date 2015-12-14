package unbeliebable

import "errors"

type Playlist struct {
	Songs []Song
}

func (m *Playlist) Add(song Song) {
	m.Songs = append(m.Songs, song)
}

func (m *Playlist) Shift() (Song, error) {
	if len(m.Songs) == 0 {
		return Song{}, errors.New("empty list")
	}
	song := m.Songs[0]
	list := m.Songs[1:]
	m.Songs = make([]Song, len(list))
	copy(m.Songs, list)
	return song, nil
}
