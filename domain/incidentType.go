package domain

type IncidentType struct {
	ID             int    `json:"id"`
	IncidentType   string `json:"incident_type"`
	Severity       string `json:"severity"`
	ResponseAction string `json:"response_action"`
	CreatedAt      string `json:"created_at"`
}
