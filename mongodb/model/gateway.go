package model


type Gateway struct {
	Id         string `bson:"_id" json:"_id"`
	// ActionId   string `bson:"actionId" json:"actionId"`
	// Campus     string `bson:"campus" json:"campus"`
	// IntranetIp string `bson:"intranetIp" json:"intranetIp"`
	InternetIp string `bson:"internetIp" json:"internetIp"`
	// Name       string `bson:"name" json:"name"`
	// Code       string `bson:"code" json:"code"`
}