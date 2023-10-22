package model

type User struct {
	UUID     string `json:"uuid" bson:"uuid"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Token struct {
	UUID  string `json:"uuid" bson:"uuid"`
	Token string `json:"token" bson:"token"`
}
