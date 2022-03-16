package options

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	srv := newService(":8000", withMaxConn(3), withTimeOut(3000))
	fmt.Println(srv)
}
