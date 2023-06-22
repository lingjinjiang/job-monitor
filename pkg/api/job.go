package api

type Job struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Namespace string      `json:"namespace"`
	Type      string      `json:"type"`
	Kind      string      `json:"kind"`
	Status    string      `json:"status"`
	Detail    interface{} `json:"detail"`
}
