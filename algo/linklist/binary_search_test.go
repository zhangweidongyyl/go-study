package linklist

import (
	"fmt"
	"testing"
)

func TestGetDaysByShipCap(t *testing.T) {
	days := getDaysByShipCap([]int{5, 3, 4}, 7)
	fmt.Printf("days is %d \r\n", days)
}
