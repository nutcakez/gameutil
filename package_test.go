package calc

import (
	"fmt"
	"math"
	"testing"

)

func TestMove(t *testing.T) {
	num := 10
	num = -num

	if num != -10 {
		t.Fatalf("nope")
	}
}

func TestMoveToPlayer(t *testing.T) {
	playerX := 0
	playerY := 0

	monsterX := 10
	monsterY := 10

	x, y := MoveToPlayer(float64(playerX), float64(playerY), float64(monsterX), float64(monsterY), 4)
	fmt.Println(x, y)
}

func TestMoveToPlayerDistanceShouldNotChangeSpeed(t *testing.T) {
	playerX := 0
	playerY := 0

	monsterX := 10
	monsterY := 10

	monsterX2 := 20
	monsterY2 := 20

	x, _ := MoveToPlayer(float64(playerX), float64(playerY), float64(monsterX), float64(monsterY), 4)
	x1, _ := MoveToPlayer(float64(playerX), float64(playerY), float64(monsterX2), float64(monsterY2), 4)
	fmt.Println("degree",math.Sinh(0.866))

	if x != x1 {
		t.Error("Expected", x, "Got:", x1)
	}

}

