package responses

type ManagerResponse struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginManagerResponse struct {
	AccessToken string          `json:"access_token"`
	Manager     ManagerResponse `json:"data"`
}
