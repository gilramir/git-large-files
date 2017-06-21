package cmd

import (
	. "gopkg.in/check.v1"
)

func (s *MySuite) TestSizeBytes(c *C) {
	var n int = -1
	n, err := parseSize("0")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 0)

	n, err = parseSize("1")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1)

	// The number can be negative, because parseSize isn't responsible
	// for checking for negatives.
	n, err = parseSize("-2")
	c.Assert(err, IsNil)
	c.Check(n, Equals, -2)

	n, err = parseSize("2.5")
	c.Assert(err, NotNil)
	c.Check(err.Error(), Equals, "The size cannot be a non-integer number of bytes")
}

func (s *MySuite) TestSizeKb(c *C) {
	var n int = -1
	n, err := parseSize("0kb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 0)

	n, err = parseSize("1kb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024)

	n, err = parseSize("1KB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024)

	n, err = parseSize("1kB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024)

	// The number can be negative, because parseSize isn't responsible
	// for checking for negatives.
	n, err = parseSize("-2kb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, -2048)

	n, err = parseSize("2.5kb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 2048+512)
}

func (s *MySuite) TestSizeMb(c *C) {
	var n int = -1
	n, err := parseSize("0mb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 0)

	n, err = parseSize("1mb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024)

	n, err = parseSize("1MB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024)

	n, err = parseSize("1mB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024)

	// The number can be negative, because parseSize isn't responsible
	// for checking for negatives.
	n, err = parseSize("-2mb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, -(1024 * 1024 * 2))

	n, err = parseSize("2.5mb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, (1024*1024*2)+(1024*1024/2))
}

func (s *MySuite) TestSizeGb(c *C) {
	var n int = -1
	n, err := parseSize("0gb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 0)

	n, err = parseSize("1gb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024*1024)

	n, err = parseSize("1GB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024*1024)

	n, err = parseSize("1gB")
	c.Assert(err, IsNil)
	c.Check(n, Equals, 1024*1024*1024)

	// The number can be negative, because parseSize isn't responsible
	// for checking for negatives.
	n, err = parseSize("-2gb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, -(1024 * 1024 * 1024 * 2))

	n, err = parseSize("2.5gb")
	c.Assert(err, IsNil)
	c.Check(n, Equals, (1024*1024*1024*2)+(1024*1024*1024/2))
}
