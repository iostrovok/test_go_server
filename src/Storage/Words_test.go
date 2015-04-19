package Storage

import (
	//"errors"
	//"Common"
	. "gopkg.in/check.v1"
	"testing"
)

func TestStorageWords(t *testing.T) {
	TestingT(t)
}

type StorageWordsTestsSuite struct{}

var _ = Suite(&StorageWordsTestsSuite{})

func _loadWords(list []string) *IndexW {
	idx := newIndexW()

	for _, word := range list {
		idx.Add(word)
	}

	return idx
}

func (s *StorageWordsTestsSuite) TestStorageWordsNew(c *C) {
	idx := newIndexW()
	c.Assert(idx, NotNil)
}

func (s *StorageWordsTestsSuite) TestStorageWordsAdd(c *C) {
	idx := newIndexW()

	m := "string"

	idx.Add(m)

	c.Assert(len(idx.cIdx), Equals, 1)
	c.Assert(len(idx.wIdx), Equals, 1)
	c.Assert(idx.wIdx[m], Equals, 0)
}

func (s *StorageWordsTestsSuite) TestStorageWordsAdd2(c *C) {

	idx := _loadWords([]string{"first", "second", "next", "second", "first", "first"})

	c.Assert(len(idx.cIdx), Equals, 3)
	c.Assert(len(idx.wIdx), Equals, 3)
	c.Assert(idx.wIdx["first"], Equals, 0)
	c.Assert(idx.wIdx["second"], Equals, 1)
	c.Assert(idx.wIdx["next"], Equals, 2)
}

func (s *StorageWordsTestsSuite) TestStorageWordsGetТор0(c *C) {
	idx := _loadWords([]string{"first", "second", "next"})
	list := idx.GetТор(0)
	c.Assert(list, NotNil)
	c.Assert(len(list), Equals, 0)
}

func (s *StorageWordsTestsSuite) TestStorageWordsGetТорShort(c *C) {
	idx := _loadWords([]string{"first", "second", "first", "next"})
	list := idx.GetТор(2)
	c.Assert(list, NotNil)
	c.Assert(len(list), Equals, 2)
	c.Assert(list[0].W, Equals, "first")
	c.Assert(list[0].N, Equals, 2)
}

func (s *StorageWordsTestsSuite) TestStorageWordsGetТорLong(c *C) {
	idx := _loadWords([]string{"first", "second", "first", "next"})
	list := idx.GetТор(5)
	c.Assert(list, NotNil)
	c.Assert(len(list), Equals, 3)
	c.Assert(list[0].W, Equals, "first")
	c.Assert(list[0].N, Equals, 2)
}
