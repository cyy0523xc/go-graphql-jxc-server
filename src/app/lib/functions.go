package lib

import (
	"io/ioutil"
)

key, e := ioutil.ReadFile("data/key")
if e != nil {
	panic(e.Error())
}

