package main

func ParseWithAI(input string) (ParseResponse, error) {

	// Temporary mock response
	return ParseResponse{
		Drugs: []Drug{
			{Name: "Paracetamol", Dosage: "500mg twice daily"},
		},
		Tests: []string{"Blood Test", "X-ray"},
		Notes: input,
	}, nil
}