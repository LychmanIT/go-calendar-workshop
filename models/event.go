package models

type Event struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Time        string   `json:"time"`
	Timezone    string   `json:"timezone"`
	Duration    int32    `json:"duration"`
	Notes       []string `json:"notes,omitempty"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
