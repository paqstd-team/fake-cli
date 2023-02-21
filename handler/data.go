package handler

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func generateData(fields map[string]string) map[string]interface{} {
	data := make(map[string]interface{})
	for key, value := range fields {
		switch value {
		// id
		case "uuid":
			data[key] = gofakeit.UUID()
		// geography
		case "city":
			data[key] = gofakeit.City()
		case "state":
			data[key] = gofakeit.State()
		case "country":
			data[key] = gofakeit.Country()
		case "latitude":
			data[key] = gofakeit.Latitude()
		case "longitude":
			data[key] = gofakeit.Longitude()
		// person
		case "name":
			data[key] = gofakeit.Name()
		case "name_prefix":
			data[key] = gofakeit.NamePrefix()
		case "name_suffix":
			data[key] = gofakeit.NameSuffix()
		case "first_name":
			data[key] = gofakeit.FirstName()
		case "last_name":
			data[key] = gofakeit.LastName()
		case "gender":
			data[key] = gofakeit.Gender()
		case "ssn":
			data[key] = gofakeit.SSN()
		case "hobby":
			data[key] = gofakeit.Hobby()
		case "email":
			data[key] = gofakeit.Email()
		case "phone":
			data[key] = gofakeit.Phone()
		case "username":
			data[key] = gofakeit.Username()
		case "password":
			data[key] = gofakeit.Password(true, true, true, true, true, 8)
		// text
		case "paragraph":
			data[key] = gofakeit.Paragraph(5, 10, 3, "\n")
		case "sentence":
			data[key] = gofakeit.Sentence(5)
		case "phrase":
			data[key] = gofakeit.Phrase()
		case "quote":
			data[key] = gofakeit.Quote()
		case "word":
			data[key] = gofakeit.Word()
		// data
		case "date":
			data[key] = gofakeit.Date()
		case "second":
			data[key] = gofakeit.Second()
		case "minute":
			data[key] = gofakeit.Minute()
		case "hour":
			data[key] = gofakeit.Hour()
		case "month":
			data[key] = gofakeit.Month()
		case "day":
			data[key] = gofakeit.Day()
		case "year":
			data[key] = gofakeit.Year()
		// internet
		case "url":
			data[key] = gofakeit.URL()
		case "domain":
			data[key] = fmt.Sprintf("%s.%s", gofakeit.DomainName(), gofakeit.DomainSuffix())
		// numbers
		case "int":
			data[key] = gofakeit.Int32()
		case "float":
			data[key] = gofakeit.Float32()
		default:
			data[key] = fmt.Sprintf("Unsupported field type: %s", value)
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
