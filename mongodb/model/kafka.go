package model

type DeleteDuplicatedRecords struct {
	ActionId string `json:"actionId"`
	Tag      string `json:"tag"`
	Action   string `json:"action"`
	Data     Data
}

type Data struct {
	Campus string `json:"campus"`
	Date   string `json:"date"`
}

type InsertCoursesToKafka struct {
	ActionId string               `json:"actionId"`
	Tag      string               `json:"tag"`
	Action   string               `json:"action"`
	Data     []SchedulesAndCamers `json:"data"`
}
