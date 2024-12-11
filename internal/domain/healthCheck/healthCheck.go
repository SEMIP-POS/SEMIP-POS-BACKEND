package healthCheck

type Health struct {
	Status      string `json:"status"`
	ServiceName string `json:"service"`
	Version     string `json:"version"`
	Database    string `json:"database"`
}
