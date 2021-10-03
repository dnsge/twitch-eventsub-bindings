package eventsub_bindings

type Subscription struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Status    string      `json:"status"`
	Cost      int         `json:"cost"`
	Condition interface{} `json:"condition"`
	CreatedAt string      `json:"created_at"`
}
