package calc

import "math"

//giving back the distance woo
func Distance(x, y, m, n float64) float64 {
	a, b := vectorTo(x, y, m, n)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func vectorTo(x, y, m, n float64) (float64, float64) {
	return x - m, y - n
}

//x and y for player
//m and n for monster
//speed the ammount of pixel to move
func MoveToPlayer(x, y, m, n, speed float64) (float64, float64) {
	distance := Distance(x, y, m, n)
	ratio := speed / distance

	xWithRatio := x / ratio
	yWithratio := y / ratio

	if x < m {
		xWithRatio = -xWithRatio
	}
	if y < n {
		yWithratio = -yWithratio
	}

	return x + xWithRatio, y + yWithratio
}

//return square of x
func Pow2(x float64) float64 {
	return math.Pow(x, 2)
}
