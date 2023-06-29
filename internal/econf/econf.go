package econf

import (
	"math/rand"
	"os"
	"strconv"
	"sync"
)

type ErpConf struct {
	JwtKey []byte
}

var (
	erpConf     *ErpConf
	erpConfOnce sync.Once
)

func (c *ErpConf) setJwtKey() {
	value := os.Getenv("JWK_KEY")
	if len(value) == 0 {
		value = strconv.Itoa(rand.Int())
	}
	c.JwtKey = []byte(value)
}

func Get() *ErpConf {
	erpConfOnce.Do(func() {
		erpConf = &ErpConf{}
		erpConf.setJwtKey()
	})
	return erpConf
}
