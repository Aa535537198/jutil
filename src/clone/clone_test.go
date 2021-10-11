package jclone

import (
	"fmt"
	"testing"
)

type School struct {
	Name string
}

func TestClone(t *testing.T) {

	sch := School{Name: "jie"}

	clone := Clone(sch)
	fmt.Println(clone)
}
