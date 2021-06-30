package service

import (
	"begin_training/mongodb/model"
	"begin_training/mongodb/sql"
	"fmt"
)

func generateRecordsForLexue(courses []model.SchedulesAndCamers) []model.SchedulesAndCamers {

	gws, _ := sql.GetGateway()

	// log.Println("输出gateway信息：", gws)
	for in1, course := range courses {
		for in2, camera := range course.Classroom.Camera {
			for _, g := range gws {
				fmt.Printf("%#v,%#v", camera.Gateway, g)
				if camera.Gateway == g.Id {
					courses[in1].Classroom.Camera[in2].GatewayMap = model.Gateway{Id: g.Id, InternetIp: g.InternetIp}
					break
				}

			}
		}

	}
	// log.Printf("输出courses信息：%#v", courses)

	return courses

}
