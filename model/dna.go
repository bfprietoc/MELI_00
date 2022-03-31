package model

type Dna struct {
	Dna []string `json:"dna" validate:"required"`
}

type ResponseDna struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
