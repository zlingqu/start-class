package model

type Schools struct { //学校
	Id       string `bson:"_id"`
	ActionId string `bson:"actionId" json:"actionId"` //actionId
	Code     string `bson:"code" json:"code"`         //code
	Name     string `bson:"name" json:"name"`         //学校名
	Type     string `bson:"type" json:"type"`         //类型
	Province string `bson:"province" json:"province"` //省份
	City     string `bson:"city" json:"city"`         //城市
	Area     string `bson:"area" json:"area"`         //区
	Address  string `bson:"address" json:"地址"`        //地址
}

type SchoolsAndCampuses struct { //学校,包含校区信息，是两个表meta-campuses和meta-schools关联之后的结果
	MySchool Schools `bson:"my-school" json:"mySchool"`
	Id       string  `bson:"_id" json:"_id"`
	ActionId string  `bson:"actionId" json:"actionId"` //actionId
	Code     string  `bson:"code" json:"code"`         //code
	Name     string  `bson:"name" json:"name"`         //学校名
	Province string  `bson:"province" json:"province"` //省份
	City     string  `bson:"city" json:"city"`         //城市
	Area     string  `bson:"area" json:"area"`         //区
	Address  string  `bson:"address" json:"地址"`        //地址
	// School   string `bson:"school" json:"school"`     //校区归属于的学校id
}

type Campuses struct { //校区
	Id       string `bson:"_id"`
	ActionId string `bson:"actionId" json:"actionId"` //actionId
	Code     string `bson:"code" json:"code"`         //code
	Name     string `bson:"name" json:"name"`         //学校名
	// Province string `bson:"province" json:"province"` //省份
	// City     string `bson:"city" json:"city"`         //城市
	// Area     string `bson:"area" json:"area"`         //区
	// Address  string `bson:"address" json:"地址"`        //地址
	School string `bson:"school" json:"school"` //校区归属于的学校id
}
