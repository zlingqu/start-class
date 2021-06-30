package service

import (
	"begin_training/mongodb/model"
	"begin_training/mongodb/sql"
	"log"
	"time"
)

var today = time.Now().Format("2006-01-02")

// func GetManuleStartClass() (courses model.SchedulesAndCamers, err error) {
func GetManuleStartClass() (courses []model.SchedulesAndCamers, err error) {
	campuses, err := sql.GetSchoolsAndCampuses()
	if len(campuses) == 0 {
		return []model.SchedulesAndCamers{}, err
	}
	for _, value := range campuses {
		log.Println("-------开始为组织【" + value.MySchool.Name + "】的校区【" + value.Name + "】生成授课记录-------")
		date, courses := getCoursesOfToday(value)
		// log.Println("曲", courses)
		if len(courses) < 1 {
			log.Println("今日没有课程,无需开课！")
			continue
		}

		deleteDuplicatedRecordsToKafka(value.Id, date) //通过发送消息到kafka的方式，删除当前的授课记录。处理哪一天就删除哪一天，业务逻辑就是这样设计的。

		if value.MySchool.Type == "cramSchool" {
			courses = generateRecordsForLexue(courses)
			// return courses, err
			insertCoursesToKafka(courses) //发往kafka
		} else {
			generateRecordsForRegularSchool(courses)
		}
	}
	log.Printf("输出courses信息：%#v", courses)
	return courses, err
}
