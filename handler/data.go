package handler

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

func generateField(value string) interface{} {
	switch value {
	// id
	case "uuid":
		return gofakeit.UUID()
	// geography
	case "city":
		return gofakeit.City()
	case "state":
		return gofakeit.State()
	case "country":
		return gofakeit.Country()
	case "latitude":
		return gofakeit.Latitude()
	case "longitude":
		return gofakeit.Longitude()
	// person
	case "name":
		return gofakeit.Name()
	case "name_prefix":
		return gofakeit.NamePrefix()
	case "name_suffix":
		return gofakeit.NameSuffix()
	case "first_name":
		return gofakeit.FirstName()
	case "last_name":
		return gofakeit.LastName()
	case "gender":
		return gofakeit.Gender()
	case "ssn":
		return gofakeit.SSN()
	case "hobby":
		return gofakeit.Hobby()
	case "email":
		return gofakeit.Email()
	case "phone":
		return gofakeit.Phone()
	case "username":
		return gofakeit.Username()
	case "password":
		return gofakeit.Password(true, true, true, true, true, 8)
	// text
	case "paragraph":
		return gofakeit.Paragraph(5, 10, 3, "\n")
	case "sentence":
		return gofakeit.Sentence(5)
	case "phrase":
		return gofakeit.Phrase()
	case "quote":
		return gofakeit.Quote()
	case "word":
		return gofakeit.Word()
	// data
	case "date":
		return gofakeit.Date()
	case "second":
		return gofakeit.Second()
	case "minute":
		return gofakeit.Minute()
	case "hour":
		return gofakeit.Hour()
	case "month":
		return gofakeit.Month()
	case "day":
		return gofakeit.Day()
	case "year":
		return gofakeit.Year()
	// internet
	case "url":
		return gofakeit.URL()
	case "domain":
		return fmt.Sprintf("%s.%s", gofakeit.DomainName(), gofakeit.DomainSuffix())
	// numbers
	case "int":
		return gofakeit.Int32()
	case "float":
		return gofakeit.Float32()
	default:
		return value
	}
}

func generateData(fields interface{}) interface{} {
	switch f := fields.(type) {
	case map[string]string:
		data := make(map[string]interface{})
		for key, value := range f {
			data[key] = generateField(value)
		}
		return data
	case map[string]interface{}:
		data := make(map[string]interface{})
		for key, value := range f {
			switch v := value.(type) {
			case string:
				data[key] = generateField(v)
			default:
				data[key] = generateData(value)
			}
		}
		return data
	case []interface{}:
		data := make([]interface{}, len(f))
		for i, item := range f {
			switch v := item.(type) {
			case string:
				data[i] = generateField(v)
			default:
				data[i] = generateData(item)
			}
		}
		return data
	default:
		return fmt.Sprintf("Unsupported type: %T", fields)
	}
}

func generateDataList(fields interface{}, page int, perPage int) []interface{} {
	if page < 1 {
		return make([]interface{}, 0)
	}

	dataList := make([]interface{}, perPage)
	for i := 0; i < perPage; i++ {
		dataList[i] = generateData(fields)
	}

	return dataList
}
