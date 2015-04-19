package Storage

/*
	Collection of words which have the same prefix.
*/
import (
	"Common"
)

// cIdx is list with reverse sort by frequency words.
// wIdx is map-index for fast access to cIdx by a word.
type IndexW struct {
	TotalWords int
	NextId     int
	LastLen    int
	cIdx       []Common.OneWord
	wIdx       map[string]int
}

func newIndexW() *IndexW {
	return &IndexW{
		cIdx: []Common.OneWord{},
		wIdx: map[string]int{},
	}
}

func (idx *IndexW) Add(m string) {
	idx.TotalWords++
	if i, find := idx.wIdx[m]; find {
		idx.cIdx[i].N++
		for i > 0 && idx.Less(i, i-1) {
			idx.Swap(i, i-1)
			i--
		}
	} else {
		w := Common.OneWord{
			N: 1,
			W: m,
		}

		if idx.LastLen <= idx.NextId {
			s := make([]Common.OneWord, 100)
			idx.cIdx = append(idx.cIdx, s...)
			idx.LastLen += 100
		}

		idx.cIdx[idx.NextId] = w
		idx.wIdx[m] = idx.NextId
		idx.NextId++
	}
}

func (idx *IndexW) GetТор(n int) []Common.OneWord {

	if n == 0 {
		return []Common.OneWord{}
	}

	if len(idx.cIdx) < n {
		return idx.cIdx
	}

	return idx.cIdx[0:n]
}

func (idx IndexW) Len() int {
	return idx.NextId
}

func (idx IndexW) Swap(i, j int) {
	idx.wIdx[idx.cIdx[i].W] = j
	idx.wIdx[idx.cIdx[j].W] = i
	idx.cIdx[i], idx.cIdx[j] = idx.cIdx[j], idx.cIdx[i]
}

func (idx IndexW) Less(i, j int) bool {
	return idx.cIdx[i].N > idx.cIdx[j].N
}
