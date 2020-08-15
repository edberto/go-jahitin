package viewmodel

type (
	LoginVM struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	RefreshVM struct {
		AccessToken string `json:"access_token"`
	}
)
