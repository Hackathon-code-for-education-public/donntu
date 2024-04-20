package domain

type Tokens struct {
	Access  string `json:"accessToken"`
	Refresh string `json:"refreshToken"`
}
