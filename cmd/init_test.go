package cmd

import (
	. "gopkg.in/check.v1"
	"log"
	"testing"
)

// Hook gocheck into the "go test" runner
func Test(t *testing.T) {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	TestingT(t)
}

type MySuite struct{}

var _ = Suite(&MySuite{})
