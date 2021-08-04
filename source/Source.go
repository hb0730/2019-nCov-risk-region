package source

type Source interface {
	Time() string
	HighRisk() []Risk
	MiddleRisk() []Risk
	Close() error
}
type Risk struct {
	Type       string   `json:"type,omitempty"`
	Province   string   `json:"province,omitempty"`
	City       string   `json:"city,omitempty"`
	County     string   `json:"county,omitempty"`
	AreaName   string   `json:"area_name,omitempty"`
	Communitys []string `json:"communitys"`
}

var Instance = NewInstanceC()

//var instanceContainer = map[string]Source{}

type InstanceC struct {
	instanceContainer map[string]Source
}

func NewInstanceC() *InstanceC {
	i := new(InstanceC)
	i.instanceContainer = map[string]Source{}
	return i
}

func (c *InstanceC) Put(name string, instance Source) {
	c.instanceContainer[name] = instance
}

func (c *InstanceC) Get(name string) Source {
	return c.instanceContainer[name]
}

func All(name string) map[string]interface{} {
	v := Instance.Get(name)
	if v == nil {
		return nil
	}
	defer v.Close()
	time := v.Time()
	hRisk := v.HighRisk()
	mRisk := v.MiddleRisk()
	return map[string]interface{}{
		"time":   time,
		"high":   hRisk,
		"middle": mRisk,
	}
}

func HighRisk(name string) map[string]interface{} {
	v := Instance.Get(name)
	if v == nil {
		return nil
	}
	defer v.Close()
	t := v.Time()
	hRisk := v.HighRisk()
	return map[string]interface{}{
		"time": t,
		"high": hRisk,
	}
}

func MiddleRisk(name string) map[string]interface{} {
	v := Instance.Get(name)
	if v == nil {
		return nil
	}
	defer v.Close()
	t := v.Time()
	mRisk := v.MiddleRisk()
	return map[string]interface{}{
		"time":   t,
		"middle": mRisk,
	}
}
