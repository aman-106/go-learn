package demo

type Movie struct {
	Name   string
	Rating Rating
}

type Rating string

const (
	R = "rated r"
	A = "uni "
)
