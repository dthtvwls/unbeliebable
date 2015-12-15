package unbeliebable

type ByVotes []Song

func (a ByVotes) Len() int {
	return len(a)
}

func (a ByVotes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByVotes) Less(i, j int) bool {
	return a[i].Score() > a[j].Score()
}
