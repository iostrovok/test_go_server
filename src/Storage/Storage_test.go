package Storage

import (
	"Common"
	. "gopkg.in/check.v1"
	"sync"
	"testing"
	"time"
)

func TestStorage(t *testing.T) {
	TestingT(t)
}

type StorageTestsSuite struct{}

var _ = Suite(&StorageTestsSuite{})

func (s *StorageTestsSuite) TestStorageNew(c *C) {
	stor, err := New()
	c.Assert(stor, NotNil)
	c.Assert(err, IsNil)
}

func checkRes(c *C, k Common.ResultMessage) {

	cw := 3
	cuw := 2
	l := 2
	if k.MyType == Common.LetterType {
		cw = 13
		cuw = 5
		l = 3
	}

	// Check CountWord
	c.Assert(k.CountWord, Equals, cw)
	// Check CountUniqWord
	c.Assert(k.CountUniqWord, Equals, cuw)
	// Check count length list
	c.Assert(len(k.Words), Equals, l)
}

func (s *StorageTestsSuite) TestStorageStart(c *C) {
	stor, _ := New()
	var WG sync.WaitGroup = sync.WaitGroup{}
	newWord := make(chan []string, 200)

	stor.Start(&WG, newWord, Common.GetFromWebClient())

	// We set all words  into the same part of storege ("ab" prefix)
	newWord <- []string{"ab", "abcde"}
	newWord <- []string{"ab", "abcd"}
	newWord <- []string{"ab", "abcd"}

	// We wait while all words are passing
	time.Sleep(1 * time.Second)

	ch := Common.SendUserRequest(3)

	k1 := <-ch
	k2 := <-ch

	checkRes(c, k1)
	checkRes(c, k2)
}

func (s *StorageTestsSuite) TestStorageStartError(c *C) {
	stor, _ := New()
	var WG sync.WaitGroup = sync.WaitGroup{}
	newWord := make(chan []string, 200)

	stor.Start(&WG, newWord, Common.GetFromWebClient())

	// Check wrong list
	newWord <- []string{"ab"}
}
