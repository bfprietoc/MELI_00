package model

type Stats struct {
	Count_mutant_dna int    `json:"count_mutant_dna"`
	Count_human_dna  int    `json:"count_human_dna"`
	Ratio            string `json:"ratio"`
}

type ResponseStats struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Stats   []Stats `json:"stats"`
}
