package entity

type Job struct {
	Id           string       `json:"id"`
	Name         string       `json:"name"`
	Desc         string       `json:"desc"`
	JobStatus    JobStatus    `json:"job_status"`
	ScheduleType ScheduleType `json:"schedule_type"`
	ScheduleExpr string       `json:"schedule_expr"`
	ExecuteType  ExecuteType  `json:"execute_type"`
	Payload      JobPayload   `json:"payload"`
	Extra        string       `json:"extra"`
	CreatedBy    string       `json:"created_by"`
	CreatedAt    int64        `json:"created_at"` // time milli
	UpdatedAt    int64        `json:"updated_at"` // time milli
}
