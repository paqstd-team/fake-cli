package handler

import (
	"github.com/brianvoe/gofakeit/v6"
)

func generateData(fields map[string]string) map[string]interface{} {
	data := make(map[string]interface{})
	for key, value := range fields {
		switch value {
		case "name":
			data[key] = gofakeit.Name()
		case "city":
			data[key] = gofakeit.City()
		case "email":
			data[key] = gofakeit.Email()
		case "uuid":
			data[key] = gofakeit.UUID()
		case "word":
			data[key] = gofakeit.Word()
		case "price":
			data[key] = gofakeit.Price(0, 1000)
		case "paragraph":
			data[key] = gofakeit.Paragraph(3, 5, 8, "\n")
		case "phrase":
			data[key] = gofakeit.Sentence(5)
		}
	}
	return data
}

func generateDataList(fields map[string]string, page int, perPage int) []map[string]interface{} {
	if page < 1 {
		return make([]map[string]interface{}, 0)
	}

	dataList := make([]map[string]interface{}, perPage)
	for i := 0; i < perPage; i++ {
		dataList[i] = generateData(fields)
	}

	return dataList
}
