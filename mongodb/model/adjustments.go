package model

// 调课表
type Adjustments struct {
	Id            string          `bson:"_id" json:"_id"`
	ActionId      string          `bson:"actionId" json:"actionId"`
	Name          string          `bson:"name" json:"name"`
	DaysOff       []string        `bson:"daysOff" json:"daysOff"`
	Campus        string          `bson:"campus" json:"campus"`
	MakeUpLessons []MakeupLessons `bson:"makeUpLessons" json:"makeUpLessons"`
}

type MakeupLessons struct {
	At  string `bson:"at" json:"at"`
	For string `bson:"for" json:"for"`
}
