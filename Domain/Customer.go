package Domain

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	DateofBirth string `json:"dateofbirth"`
	Status      string `json:"status"`
}
