package api

// Log of an operation of a deployment
type Log struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Title   string `json:"title"`
	Payload struct {
		Branch string `json:"branch"`
		Commit struct {
			Author  string `json:"author"`
			Message string `json:"message"`
			SHA     string `json:"sha"`
		}
	}
}
