package model

type User struct {
	UserId       string `json:"user_id" bson:"user_id"`
	Email        string `json:"email" bson:"email"`
	Lock         bool   `json:"lock" bson:"lock"`
	UserName     string `json:"username" bson:"username"`
	Avatar       []byte `json:"avatar" bson:"avatar"`
	Pass         string `json:"pass" bson:"password"`
	Try          int    `json:"try" bson:"try"`
	UnlockTime   string `json:"unlockTime" bson:"unlock_time"`
	Token        string `json:"token" bson:"-"`
	NationalId   string `json:"national_id" bson:"national_id"`
	DataeOfBirth string `json:"date_of_birth" bson:"date_of_birth"`
	IsVerified   string `json:"is_verified" bson:"is_verified"`
	CreateAt     string `json:"create_at" bson:"create_at"`
	UpdateAt     string `json:"update_at" bson:"update_at"`
	//tableName     struct{} `bson:"user"`
}
