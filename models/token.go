package models

type AccessDetails struct {
	AccessUUID string
	UserID     string
}

type TokenDetails struct {
	AccessToken string
	AccessUuid  string
	AtExpires   int64
}
