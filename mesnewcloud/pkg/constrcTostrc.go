package pkg

import (
	"reflect"
	"strings"
	"time"
)

func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		//fmt.Println("type", obj1.Field(i).Type.String())
		if !strings.Contains(obj1.Field(i).Type.String(), "time") && obj1.Field(i).Type.String() != "int" {
			data[obj1.Field(i).Name] = StructToMap(obj2.Field(i).Interface())
			//fmt.Println("dadada", data)
		} else if strings.Contains(obj1.Field(i).Type.String(), "time") {
			t := obj2.Field(i).Interface().(time.Time)
			data[obj1.Field(i).Name] = t.Format("2006-01-02 15:04:05")
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}
func StructToMap(obj interface{}) interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var name string
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Type.String() != "bool" && !strings.Contains(obj1.Field(i).Type.String(), "Time") {
			//	fmt.Println("+++", obj1.Field(i).Type.String())
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
			name = obj1.Field(i).Name
		} else if strings.Contains(obj1.Field(i).Type.String(), "Time") {
			t := obj2.Field(i).Interface().(time.Time)
			name = obj1.Field(i).Name
			data[obj1.Field(i).Name] = t.Format("2006-01-02 15:04:05")
		}
	}

	return data[name]
}
