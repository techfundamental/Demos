package models

type Plan struct {
	Version string `json:"format_version"`
	Planned struct {
		Outputs struct {
			Name struct {
				Value string `json:"value"`
			} `json:"name"`
		} `json:"outputs"`
		Modules struct {
			Deployments []Resource `json:"resources"`
		} `json:"root_module"`
	} `json:"planned_values"`
}
