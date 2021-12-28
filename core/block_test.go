package core

import (
	"testing"
)

func TestBlockPrint(t *testing.T) {
	block1 := Block{Height: 432, Width: 432}

	str := block1.String()

	if str != "height is 432 and width is 432" {
		t.Errorf("Block{Height: 432, Width: 432} = %s; want height is 432 and width is 432", str)
	}
}
