package model

type SchedulesAndCamers struct { //课程表
	Id        string    `bson:"_id" json:"_id"`
	ActionId  string    `bson:"actionId" json:"actionId"`
	Type      string    `bson:"type" json:"type"`
	From      string    `bson:"from" json:"from"`
	To        string    `bson:"to" json:"to"`
	Class     Class     `bson:"class" json:"class"`
	Teacher   Teacher   `bson:"teacher" json:"teacher"`
	Subject   Subject   `bson:"subject" json:"subject"`
	Classroom Classroom `bson:"classroom" json:"classroom"`
	Schedule  Schedule  `bson:"schedule" json:"schedule"`
	Campus    string    `bson:"campus" json:"campus"`
	StartTime string    `bson:"startTime" json:"startTime"` //来自于schedule
	EndTime   string    `bson:"endTime" json:"endTime"`     //来自于schedule
	Weekday   int       `bson:"weekday" json:"weekday"`     //来自于schedule
	Date      string    `bson:"date" json:"date"`           //新增字段，不来自于数据库
	Week      int       `bson:"week" json:"week"`           //新增字段，不来自于数据库
}

type Schedule struct {
	Weekday   int    `bson:"weekday" json:"weekday"`
	StartTime string `bson:"startTime" json:"startTime"`
	EndTime   string `bson:"etartTime" json:"endTime"`
}

type Classroom struct {
	Id     string    `bson:"_id" json:"_id"`
	Code   string    `bson:"code" json:"code"`
	Name   string    `bson:"name" json:"name"`
	Camera []Cameras `bson:"cameras" json:"cameras"`
}

type Subject struct {
	Id     string `bson:"_id" json:"_id"`
	Type   string `bson:"type" json:"type"`
	Weight int    `bson:"weight" json:"weight"`
}

type Teacher struct {
	Id   string `bson:"_id" json:"_id"`
	Code string `bson:"code" json:"code"`
	Name string `bson:"name" json:"name"`
}

type Class struct {
	Id               string `bson:"_id" json:"_id"`
	Name             string `bson:"name" json:"name"`
	Code             string `bson:"code" json:"code"`
	Weight           int    `bson:"weight" json:"weight"`
	ClassTeacher     string `bson:"classTeacher" json:"classTeacher"`
	NumberOfStudents int    `bson:"numberOfStudents" json:"numberOfStudents"`
	Grade            Grade
}

type Grade struct {
	Id     string `bson:"_id" json:"_id"`
	Name   string `bson:"name" json:"name"`
	Code   string `bson:"code" json:"code"`
	Weight int    `bson:"weight" json:"weight"`
}

type Cameras struct {
	GatewayMap    Gateway `bson:"gatewayMap" json:"gateway"`
	Id            string  `bson:"_id" json:"_id"`
	ActionId      string  `bson:"actionId" json:"-"`
	CameraUrl     string  `bson:"url" json:"cameraUrl"`
	Gateway       string  `bson:"gateway" json:"-"`
	Classroom     string  `bson:"classroom" json:"-"`
	Position      string  `bson:"position" json:"position"`
	Campus        string  `bson:"campus" json:"-"`
	LowBitRateUrl string  `bson:"lowBitRateUrl" json:"lowBitRateUrl"`
}

type Labels struct {
	CoursesProgress CoursesProgress `bson:"coursesProgress" json:"coursesProgress"`
}

type CoursesProgress struct {
	Total int `bson:"total" json:"total"`
}
