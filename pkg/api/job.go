package api

type Overview struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Type      string `json:"type"`
	Kind      string `json:"kind"`
	Status    string `json:"status"`
}

type Job struct {
	Overview
	Detail interface{} `json:"detail"`
}
