package models

type Validation struct {
	IsValid bool `json:"valid"`
	ErrorCount int `json:"error_count"`
	WarningCount int `json:"warning_count"`
}
