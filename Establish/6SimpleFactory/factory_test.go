package _SimpleFactory

import (
	"testing"
)

func TestNewRestaurant(t *testing.T) {
	NewRestaurant("d").GetFood()
	NewRestaurant("q").GetFood()
}
