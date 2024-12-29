package entities

type Session struct {
	UserID      int    `json:"userId"`
	AccessToken string `json:"accessToken"`
}

func NewSession(userId int, accessToken string) *Session {
	return &Session{
		UserID:      userId,
		AccessToken: accessToken,
	}
}
