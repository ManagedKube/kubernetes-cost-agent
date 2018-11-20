package price

type Instance struct {
	Name       string     `json:"name"`
	Memory     int64      `json:"memory"`
	Cpu        int        `json:"cpu"`
	HourlyCost HourlyCost `json:"hourlyCost"`
}

type HourlyCost struct {
	OnDemand   float64 `json:"onDemand"`
	ReduceCost float64 `json:"reduceCost"`
}

type Instances struct {
	Instance []Instance `json:"instances"`
}
