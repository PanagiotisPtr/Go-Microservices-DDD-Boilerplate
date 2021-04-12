package dto

type RefreshTokenDto struct {
	Uuid       string
	UserUuid   string
	Expiration int64
}
