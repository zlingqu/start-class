package service

import (
	"begin_training/mongodb/model"
	"begin_training/mongodb/sql"
	"begin_training/tools"
	"fmt"
)

//调课
func getScheduleDateAfterAdjusting(adjustment model.Adjustments, shoolType string) string {

	if tools.Find(adjustment.DaysOff, today) { //今天是假期
		return ""
	}
	for _, i := range adjustment.MakeUpLessons { //调课到了今天，返回for日期值，为后者生成课表
		if i.At == today {
			return i.For
		}
	}
	return today
}

// 排查调课、排查已经结束的班级后的课程表
func getCoursesOfToday(campus model.SchoolsAndCampuses) (string, []model.SchedulesAndCamers) { //

	// var adjustmentData []model.Adjustments
	// adjustmentData, _ = sql.GetAdjustments(campus.Id) //查询调课表
	// if len(adjustmentData) == 0 {                     //没有调课记录，返回today
	// 	log.Printf("没有调课记录,只完成今日的即可！")
	// 	// return today, []model.SchedulesAndCamers{}
	// }
	// date := getScheduleDateAfterAdjusting(adjustmentData[0], campus.MySchool.Type) //这里可能逻辑有问题，只会判定数组中的第一条。培佳说，理论上调课记录只能有一条

	// if date == "" { //放假不用上课
	// 	log.Printf("今天放假，无需执行操作！")
	// 	return date, []model.SchedulesAndCamers{}
	// }
	courses, _ := sql.GetSchedulesAndCameras(today, campus.Id) //查找课程表和校区的对应信息
	// var courses []model.SchedulesAndCamers
	fmt.Println("该小区找到的课程数量为", len(courses))

	classData, _ := sql.GetTodayClass(today, campus.Id) //获取已经结束的班级信息，earlyTerminationDate=true
	// fmt.Println("找到的course数量为", len(courses))

	for i := 0; i < len(courses); i++ {
		for _, value := range classData {
			if courses[i].Class.Id == value.Id {
				courses[i] = model.SchedulesAndCamers{} //班级有有效期，有效班级已经过期，班级对应的课程也就不需要了。
			}
		}
	}
	fmt.Println("剔除过期课程", len(courses))
	return today, courses
}
