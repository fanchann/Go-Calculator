package repository

type Calculator struct{}

func NewCalculator(c *Calculator) Calculate {
	return c
}

func (c *Calculator) Add(a int, b int) int {
	return a + b
}

func (c *Calculator) Multiply(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a * b
}

func (c *Calculator) Divide(a int, b int) float64 {
	if a == 0 || b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}

func (c *Calculator) Less(a int, b int) float64 {
	return float64(a) - float64(b)
}
