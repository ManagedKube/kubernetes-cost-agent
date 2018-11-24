package persistentVolume

type PersistentVolume struct {
	Name                 string
	Capacity             int64
	VolumeName           string
	StatusPhase          string
	SpecStorageClassName string
	Claim                PersistentVolumeClaim
	CostPerGbMonth       float64
	CostPerGbHour        float64
}

type PersistentVolumeList struct {
	PersistentVolume []PersistentVolume
}

type PersistentVolumeClaim struct {
	Name      string
	Namespace string
	Kind      string
}
