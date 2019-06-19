package types

//go:generate jsonenums -type=AdType

type AdType int

const (
	AdCarousel AdType = iota
	AdHomeRecommendation
)
