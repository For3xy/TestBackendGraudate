package domain

type Situation struct {
	ID             int    `json:"id"`
	IncidentTypeID int    `json:"incident_type_id"`
	CameraID       int    `json:"camera_id"`
	Description    string `json:"description"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	CreatedAt      string `json:"created_at"`
}
