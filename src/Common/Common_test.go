package Common

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestNew(t *testing.T) {
	TestingT(t)
}

type CommonTestsSuite struct{}

var _ = Suite(&CommonTestsSuite{})

func (s *CommonTestsSuite) TestCommonGetFromWebClient(c *C) {
	c.Assert(fromWebClient, IsNil)
	res := GetFromWebClient()
	c.Assert(res, NotNil)
	c.Assert(fromWebClient, NotNil)
}

func (s *CommonTestsSuite) TestCommonSendUserRequest(c *C) {
	res := SendUserRequest(11)
	c.Assert(res, NotNil)

	m := <-fromWebClient
	c.Assert(m.N, Equals, 11)
	c.Assert(m.Ch, NotNil)
}

func (s *CommonTestsSuite) TestCommonSortAndCutEmpty(c *C) {

	list := []OneWord{}

	res := SortAndCut(3, list)
	c.Assert(res, NotNil)
	c.Assert(len(res), Equals, 0)
}

func (s *CommonTestsSuite) TestCommonSortAndCutZero(c *C) {

	list := []OneWord{
		OneWord{"e", 12},
		OneWord{"d", 1},
	}

	res := SortAndCut(0, list)
	c.Assert(res, NotNil)
	c.Assert(len(res), Equals, 0)
}

func (s *CommonTestsSuite) TestCommonSortAndCut(c *C) {

	list := []OneWord{
		OneWord{"e", 12},
		OneWord{"a", 11},
		OneWord{"b", 14},
		OneWord{"c", 15},
		OneWord{"d", 1},
	}

	res := SortAndCut(3, list)
	c.Assert(res, NotNil)
	c.Assert(res[0], Equals, "c")
}
