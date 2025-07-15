package model

type UserData struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Age     *int   `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type UserDataInput struct {
	Name    string `json:"name"`
	Age     *int   `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
