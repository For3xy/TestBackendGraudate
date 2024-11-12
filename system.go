package testbackendGraudate

type Camera struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

type Situation struct {
	ID             int    `json:"id"`
	IncidentTypeID int    `json:"incident_type_id"`
	CameraID       int    `json:"camera_id"`
	Description    string `json:"description"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	CreatedAt      string `json:"created_at"`
}

type IncidentType struct {
	ID             int    `json:"id"`
	IncidentType   string `json:"incident_type"`
	Severity       string `json:"severity"`
	ResponseAction string `json:"response_action"`
	CreatedAt      string `json:"created_at"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
}

type Notification struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	SituationID int    `json:"situation_id"`
	Message     string `json:"message"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	SentAt      string `json:"sent_at"`
}

type UserSituation struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	SituationID int    `json:"situation_id"`
	AssignedAt  string `json:"assigned_at"`
}
