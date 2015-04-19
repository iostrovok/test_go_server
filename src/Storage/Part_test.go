package Storage

import (
	"Common"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func TesStoragePart(t *testing.T) {
	TestingT(t)
}

type StoragePartTestsSuite struct{}

var _ = Suite(&StoragePartTestsSuite{})

func (s *StoragePartTestsSuite) TestStoragePartNew(c *C) {
	res := newPart(Common.WordType, "ab")
	c.Assert(res, NotNil)
}

func (s *StoragePartTestsSuite) TestStoragePartAdd(c *C) {
	part := newPart(Common.WordType, "ab")
	part.Add("adcd")
	c.Assert(part, NotNil)
}

func (s *StoragePartTestsSuite) TestStoragePartTop(c *C) {
	part := newPart(Common.WordType, "ab", true)
	part.Add("abcde")
	part.Add("abcd")
	part.Add("abcd")

	// We wait while all words are passing
	time.Sleep(1 * time.Second)

	ch := make(chan Common.ResultMessage, 10)
	m := Common.UserRequest{
		N:  3,
		Ch: ch,
	}

	part.GetTop(1, m)
	k := <-ch

	c.Assert(k.CountWord, Equals, 3)
	c.Assert(k.CountUniqWord, Equals, 2)
	c.Assert(len(k.Words), Equals, 2)
}
