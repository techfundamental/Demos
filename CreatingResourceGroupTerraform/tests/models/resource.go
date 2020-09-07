package models

type Resource struct {
	Address string `json:"address"`
	Mode string `json:"mode"`
	Type string `json:"type"`
	Name string `json:"name"`
	Provider string `json:"provider_name"`
	Inputs map[string]interface{} `json:"values"`
}
