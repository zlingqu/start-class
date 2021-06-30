package tools

import (
	"reflect"
)

func Find(slice []string, val string) bool { //判断切片中是否包含某个元素
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
