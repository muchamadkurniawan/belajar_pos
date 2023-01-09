package helper

import (
	"belajar_pos/model/web"
	"encoding/json"
)

func StructToMap(data web.KasirResponse) map[string]interface{} {
	var myMap map[string]interface{}
	dataConv, _ := json.Marshal(data)
	json.Unmarshal(dataConv, &myMap)
	return myMap
}

func StructSliceToMap(datas []web.KasirResponse) []map[string]interface{} {
	var dataMAPs []map[string]interface{}
	for index, _ := range datas {
		data := datas[index]
		var myMap map[string]interface{}
		dataConv, _ := json.Marshal(data)
		json.Unmarshal(dataConv, &myMap)
		dataMAPs = append(dataMAPs, myMap)
	}
	return dataMAPs
}
