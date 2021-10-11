package jtime

import (
	"fmt"
	"testing"
	"time"
)

func TestAddYear(t *testing.T) {
	now := time.Now()

	AddTime(&now, 100, 1, 1)

	fmt.Println(now)
}
