package persistentVolume

type PersistentVolume struct {
	Name                 string                `json:"name"`
	Capacity             int64                 `json:"capacity"`
	VolumeName           string                `json:"volumeName"`
	StatusPhase          string                `json:"statusPhase"`
	SpecStorageClassName string                `json:"specStorageClassName"`
	Claim                PersistentVolumeClaim `json:"claim"`
	CostPerGbMonth       float64               `json:"costPerGbMonth"`
	CostPerGbHour        float64               `json:"costPerGbHour"`
}

type PersistentVolumeList struct {
	PersistentVolume []PersistentVolume `json:"persistentVolume"`
}

type PersistentVolumeClaim struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Kind      string `json:"kind"`
}
