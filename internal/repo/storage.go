package repo

type StorageRepo interface {
	UpdateMoistureSettings(data []byte) error
	GetMoistureSettings() ([]byte, error)
}
