package main

type ParseRequest struct {
	Text string `json:"text"`
}

type Drug struct {
	Name   string `json:"name"`
	Dosage string `json:"dosage"`
}

type ParseResponse struct {
	Drugs []Drug  `json:"drugs"`
	Tests []string `json:"tests"`
	Notes string   `json:"notes"`
}

type BillResponse struct {
	Total int `json:"total"`
}