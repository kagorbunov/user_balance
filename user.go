package avitoAPI



type User struct{
	Id int `json:"-"`
	Name string `json:"name"`
	Balance float64 `json:"balance"`
}