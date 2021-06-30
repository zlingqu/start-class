package service

import (
	"begin_training/kafka/producer"
	"begin_training/mongodb/model"
	"encoding/json"
	"log"
)

func deleteDuplicatedRecordsToKafka(campusId, date string) {
	msg_data := &model.Data{
		Campus: campusId,
		Date:   date,
	}
	msg := &model.DeleteDuplicatedRecords{
		ActionId: "123",
		Tag:      "meta-records",
		Action:   "del",
		Data:     *msg_data,
	}

	b, _ := json.Marshal(msg)
	producer.ProductMessToKafka("XMC2-CUD-META-SERVICE", string(b))
	log.Printf("这里会发送到kafka的topic(meta-records),清空当天的授课记录，%T", b)
}

func insertCoursesToKafka(courses []model.SchedulesAndCamers) {

	msg := &model.InsertCoursesToKafka{
		ActionId: "456",
		Tag:      "meta-records",
		Action:   "insert",
		Data:     courses,
	}

	b, _ := json.Marshal(msg)
	producer.ProductMessToKafka("XMC2-CUD-META-SERVICE", string(b))
	log.Printf("这里会发送到kafka的topic(meta-records),课程记录，%T", b)
}
