package socks

import (
	"testing"
)

func TestGenDefaultHttpResponse(t *testing.T) {
	t.Parallel()
	t.Log(GenDefaultHttpResponse("1.2.3.4"))
}