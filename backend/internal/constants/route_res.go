package constants

type Leg struct {
	Steps    []string `json:"steps"`
	Summary  string   `json:"summary"`
	Weight   float64  `json:"weight"`
	Duration float64  `json:"duration"`
	Distance float64  `json:"distance"`
}
type Route struct {
	Legs       []Leg   `json:"legs"`
	WeightName string  `json:"weight_name"`
	Weight     float64 `json:"weight"`
	Duration   float64 `json:"duration"`
	Distance   float64 `json:"distance"`
}
type RouteResp struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}
