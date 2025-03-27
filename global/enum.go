package global

type ImportanceType int

const (
	Light ImportanceType = iota
	Normal
	Important
	Critical
)
