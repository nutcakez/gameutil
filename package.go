package calc

import "math"

//giving back the distance woo
func Distance(x, y, m, n float64) float64 {
	a, b := vectorTo(x, y, m, n)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func vectorTo(x, y, m, n float64) (float64, float64) {
	return m - x, n -y 
}

//x and y for player
//m and n for monster
//speed the ammount of pixel to move
func MoveToPlayer(x, y, m, n, speed float64) (float64, float64) {
	distance := Distance(x, y, m, n)

	b, a := vectorTo(x,y,m,n)

	ratio := speed/distance

	return x+(a*ratio), y+(b*ratio)
}

func MoveToPlayer32(x, y, m, n, speed float32) (float64, float64) {
	return MoveToPlayer(float64(x), float64(y), float64(m), float64(n), float64(speed))
}

//return square of x
func Pow2(x float64) float64 {
	return math.Pow(x, 2)
}

func CircleCircleCollision(x, y, radius, m, n, radius2 float64) bool {
	distance := Distance(x, y, m, n)
	return distance <= m+n
}

func InvSin(x float64) float64 {
	return math.Asin(x) * (180 / math.Pi)
}

func InvCos(x float64) float64 {
	return math.Acos(x) * (180/math.Pi)
}