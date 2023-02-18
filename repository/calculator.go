package repository

type Calculate interface {
	Add(a, b int) int
	Multiply(a, b int) int
	Divide(a, b int) float64
	Less(a, b int) float64
}
