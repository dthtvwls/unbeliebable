package unbeliebable

import (
	"errors"
	"net"
	"time"
)

type Song struct {
	IP               net.IP
	Time             time.Time
	ID, Name, Artist string
	Votes            []Vote
}

func (m *Song) Vote(vote Vote) error {
	for i := range m.Votes {
		if m.Votes[i].IP.Equal(vote.IP) {
			return errors.New("already voted")
		}
	}
	m.Votes = append(m.Votes, vote)
	return nil
}

func (m *Song) Disqualifies() bool {
	downvotes := 0
	for i := range m.Votes {
		if m.Votes[i].Value < 0 {
			downvotes += m.Votes[i].Value
		}
		if downvotes <= -3 {
			return true
		}
	}
	return false
}

func (m *Song) Score() int {
	sum := 0
	for i := range m.Votes {
		sum += m.Votes[i].Value
	}
	return sum
}
