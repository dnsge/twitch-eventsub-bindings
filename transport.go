package eventsub_bindings

type Request struct {
	Type      string      `json:"type"`
	Version   string      `json:"version"`
	Condition interface{} `json:"condition"`
	Transport Transport   `json:"transport"`
}

type Response struct {
	Subscription Subscription `json:"subscription"`
	Event        interface{}  `json:"event"`
}

type Transport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}
