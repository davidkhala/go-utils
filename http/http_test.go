package http

import (
	"fmt"
	"github.com/davidkhala/goutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	resp := Get("http://www.google.com", nil)
	fmt.Println(string(resp.BodyBytes()))
}
func TestTimeout(t *testing.T) {
	var invalidUrl = "http://abc.google.com"
	var handler = func(errString string, params ...interface{}) (success bool) {
		return assert.Regexp(t, "^Get "+invalidUrl+": dial tcp: lookup ", errString)
	}
	defer goutils.Deferred(handler)
	Get(invalidUrl, nil)

}
