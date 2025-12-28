package entity

type ScheduleType int64

const (
	ScheduleTypeImmediate ScheduleType = iota
	ScheduleTypeDelayed
	ScheduleTypePeriodic
)

type ExecuteType int64

const (
	ExecuteTypeRpc ExecuteType = iota
	ExecuteTypeHttp
	ExecuteTypeGolang
	ExecuteTypePython
	ExecuteTypeShell
)

type JobStatus int64

const (
	JobStatusPending JobStatus = iota
	JobStatusActive
	JobStatusPaused
	JobStatusDisabled
	JobStatusArchived
)
