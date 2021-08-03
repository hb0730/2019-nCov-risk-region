package source

type Source interface {
	Time() (string, error)
	HighRisk() []Risk
	MiddleRisk() []Risk
	Close() error
}
type Risk struct {
	Type       string
	Province   string
	City       string
	County     string
	AreaName   string
	communitys []string
}
