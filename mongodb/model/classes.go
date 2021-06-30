package model

type Classes struct { //课时
	Id       string `bson:"_id"`
	ActionId string `bson:"actionId" json:"actionId"`
	Code      string `bson:"code" json:"code"`
	Name      string `bson:"name" json:"name"`
	Weight    int    `bson:"weight" json:"weight"`
	From      string `bson:"from" json:"from"`
	To        string `bson:"to" json:"to"`
	Class     string `bson:"class" json:"class"`
	Grade     string `bson:"grade" json:"grade"`
	Campus    string `bson:"campus" json:"campus"`
	Classroom string `bson:"classroom" json:"classroom"`
	EarlyTerminationDate string `bson:"earlyTerminationDate" json:"earlyTerminationDate"` //提前结束的班级
}
