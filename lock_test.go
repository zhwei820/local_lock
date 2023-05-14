package local_lock

import (
	"fmt"
	"testing"
)

func TestLock(t *testing.T) {
	unlock := Lock("amm_order_0001") // lock by key
	fmt.Println("1")
	unlock()
	unlock2 := Lock("amm_order_0001") // lock by key
	fmt.Println("2")
	defer unlock2()
}
