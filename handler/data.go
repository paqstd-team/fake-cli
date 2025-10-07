package handler

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
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
	case "address":
		return gofakeit.Address().Address
	case "street":
		return gofakeit.Address().Street
	case "zip":
		return gofakeit.Address().Zip
	case "postal_code":
		return gofakeit.Address().Zip
	case "timezone":
		return gofakeit.TimeZone()
	case "timezone_abbr":
		return gofakeit.TimeZoneAbv()
	case "timezone_full":
		return gofakeit.TimeZoneFull()
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
	case "company":
		return gofakeit.Company()
	case "job_title":
		return gofakeit.JobTitle()
	case "job_descriptor":
		return gofakeit.JobDescriptor()
	case "job_level":
		return gofakeit.JobLevel()
	case "bs":
		return gofakeit.BS()
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
	case "datetime":
		return gofakeit.Date()
	case "time":
		return gofakeit.Date()
	case "weekday":
		return gofakeit.WeekDay()
	case "month_string":
		return gofakeit.MonthString()
	case "price":
		return gofakeit.Price(0.50, 1000.00)
	case "currency":
		return gofakeit.CurrencyShort()
	case "currency_long":
		return gofakeit.CurrencyLong()
	case "currency_code":
		return gofakeit.CurrencyShort()
	case "credit_card":
		return gofakeit.CreditCardNumber(nil)
	case "credit_card_cvv":
		return gofakeit.CreditCardCvv()
	case "credit_card_exp":
		return gofakeit.CreditCardExp()
	case "credit_card_type":
		return gofakeit.CreditCardType()
	case "cvv":
		return gofakeit.CreditCardCvv()
	case "cvc":
		return gofakeit.CreditCardCvv()
	case "expiry":
		return gofakeit.CreditCardExp()
	case "expiration":
		return gofakeit.CreditCardExp()
	// Banking
	case "bank_name":
		return gofakeit.BankName()
	case "bank_type":
		return gofakeit.BankType()
	case "ein":
		return gofakeit.EIN()
	case "ach_account":
		return gofakeit.AchAccount()
	case "ach_routing":
		return gofakeit.AchRouting()
	// internet
	case "url":
		return gofakeit.URL()
	case "domain":
		return fmt.Sprintf("%s.%s", gofakeit.DomainName(), gofakeit.DomainSuffix())
	case "domain_name":
		return gofakeit.DomainName()
	case "domain_suffix":
		return gofakeit.DomainSuffix()
	case "ip":
		return gofakeit.IPv4Address()
	case "ipv4":
		return gofakeit.IPv4Address()
	case "ipv6":
		return gofakeit.IPv6Address()
	case "mac_address":
		return gofakeit.MacAddress()
	case "http_method":
		return gofakeit.HTTPMethod()
	case "http_status_code":
		return gofakeit.HTTPStatusCode()
	case "http_status_simple":
		return gofakeit.HTTPStatusCodeSimple()
	case "user_agent":
		return gofakeit.UserAgent()
	case "chrome_user_agent":
		return gofakeit.ChromeUserAgent()
	case "firefox_user_agent":
		return gofakeit.FirefoxUserAgent()
	case "safari_user_agent":
		return gofakeit.SafariUserAgent()
	case "opera_user_agent":
		return gofakeit.OperaUserAgent()
	// products
	case "product_name":
		return gofakeit.ProductName()
	case "product_category":
		return gofakeit.ProductCategory()
	case "product_description":
		return gofakeit.ProductDescription()
	case "product_feature":
		return gofakeit.ProductFeature()
	case "product_material":
		return gofakeit.ProductMaterial()
	case "product_upc":
		return gofakeit.ProductUPC()
	case "product_audience":
		return gofakeit.ProductAudience()
	case "product_benefit":
		return gofakeit.ProductBenefit()
	case "product_dimension":
		return gofakeit.ProductDimension()
	case "product_isbn":
		return gofakeit.ProductISBN(nil)
	case "product_suffix":
		return gofakeit.ProductSuffix()
	case "product_use_case":
		return gofakeit.ProductUseCase()
	case "brand":
		return gofakeit.CarMaker()
	case "color":
		return gofakeit.Color()
	case "hex_color":
		return gofakeit.HexColor()
	case "rgb_color":
		return gofakeit.RGBColor()
	case "safe_color":
		return gofakeit.SafeColor()
	// animals
	case "animal":
		return gofakeit.Animal()
	case "animal_type":
		return gofakeit.AnimalType()
	case "bird":
		return gofakeit.Bird()
	case "cat":
		return gofakeit.Cat()
	case "dog":
		return gofakeit.Dog()
	case "farm_animal":
		return gofakeit.FarmAnimal()
	case "pet_name":
		return gofakeit.PetName()
	// food
	case "breakfast":
		return gofakeit.Breakfast()
	case "lunch":
		return gofakeit.Lunch()
	case "dinner":
		return gofakeit.Dinner()
	case "snack":
		return gofakeit.Snack()
	case "dessert":
		return gofakeit.Dessert()
	case "drink":
		return gofakeit.Drink()
	case "fruit":
		return gofakeit.Fruit()
	case "vegetable":
		return gofakeit.Vegetable()
	// beer
	case "beer_name":
		return gofakeit.BeerName()
	case "beer_style":
		return gofakeit.BeerStyle()
	case "beer_hop":
		return gofakeit.BeerHop()
	case "beer_malt":
		return gofakeit.BeerMalt()
	case "beer_yeast":
		return gofakeit.BeerYeast()
	case "beer_alcohol":
		return gofakeit.BeerAlcohol()
	case "beer_blg":
		return gofakeit.BeerBlg()
	case "beer_ibu":
		return gofakeit.BeerIbu()
	// cars
	case "car_maker":
		return gofakeit.CarMaker()
	case "car_model":
		return gofakeit.CarModel()
	case "car_type":
		return gofakeit.CarType()
	case "car_fuel_type":
		return gofakeit.CarFuelType()
	case "car_transmission_type":
		return gofakeit.CarTransmissionType()
	// movies
	case "movie_name":
		return gofakeit.MovieName()
	case "movie_genre":
		return gofakeit.MovieGenre()
	// books
	case "book_title":
		return gofakeit.BookTitle()
	case "book_author":
		return gofakeit.BookAuthor()
	case "book_genre":
		return gofakeit.BookGenre()
	// music
	case "song":
		return gofakeit.Song()
	case "song_artist":
		return gofakeit.SongArtist()
	case "song_genre":
		return gofakeit.SongGenre()
	case "song_name":
		return gofakeit.SongName()
	// apps
	case "app_name":
		return gofakeit.AppName()
	case "app_author":
		return gofakeit.AppAuthor()
	case "app_version":
		return gofakeit.AppVersion()
	// numbers
	case "int":
		return gofakeit.Int32()
	case "int8":
		return gofakeit.Int8()
	case "int16":
		return gofakeit.Int16()
	case "int32":
		return gofakeit.Int32()
	case "int64":
		return gofakeit.Int64()
	case "uint8":
		return gofakeit.Uint8()
	case "uint16":
		return gofakeit.Uint16()
	case "uint32":
		return gofakeit.Uint32()
	case "uint64":
		return gofakeit.Uint64()
	case "float":
		return gofakeit.Float32()
	case "float32":
		return gofakeit.Float32()
	case "float64":
		return gofakeit.Float64()
	case "bool":
		return gofakeit.Bool()
	case "number":
		return gofakeit.Number(1, 100)
	case "int_n":
		return gofakeit.IntN(10)
	case "uint_n":
		return gofakeit.UintN(10)
	case "float32_range":
		return gofakeit.Float32Range(0.0, 100.0)
	case "float64_range":
		return gofakeit.Float64Range(0.0, 100.0)
	case "digit":
		return gofakeit.Digit()
	case "digit_n":
		return gofakeit.DigitN(3)
	case "letter":
		return gofakeit.Letter()
	case "letter_n":
		return gofakeit.LetterN(5)
	case "vowel":
		return gofakeit.Vowel()
	// crypto
	case "bitcoin_address":
		return gofakeit.BitcoinAddress()
	case "bitcoin_private_key":
		return gofakeit.BitcoinPrivateKey()
	// other
	case "emoji":
		return gofakeit.Emoji()
	case "emoji_alias":
		return gofakeit.EmojiAlias()
	case "emoji_category":
		return gofakeit.EmojiCategory()
	case "emoji_description":
		return gofakeit.EmojiDescription()
	case "emoji_tag":
		return gofakeit.EmojiTag()
	case "gamertag":
		return gofakeit.Gamertag()
	// minecraft
	case "minecraft_animal":
		return gofakeit.MinecraftAnimal()
	case "minecraft_armor_part":
		return gofakeit.MinecraftArmorPart()
	case "minecraft_armor_tier":
		return gofakeit.MinecraftArmorTier()
	case "minecraft_biome":
		return gofakeit.MinecraftBiome()
	case "minecraft_dye":
		return gofakeit.MinecraftDye()
	case "minecraft_food":
		return gofakeit.MinecraftFood()
	case "minecraft_mob_boss":
		return gofakeit.MinecraftMobBoss()
	case "minecraft_mob_hostile":
		return gofakeit.MinecraftMobHostile()
	case "minecraft_mob_neutral":
		return gofakeit.MinecraftMobNeutral()
	case "minecraft_mob_passive":
		return gofakeit.MinecraftMobPassive()
	case "minecraft_ore":
		return gofakeit.MinecraftOre()
	case "minecraft_tool":
		return gofakeit.MinecraftTool()
	case "minecraft_villager_job":
		return gofakeit.MinecraftVillagerJob()
	case "minecraft_villager_level":
		return gofakeit.MinecraftVillagerLevel()
	case "minecraft_villager_station":
		return gofakeit.MinecraftVillagerStation()
	case "minecraft_weapon":
		return gofakeit.MinecraftWeapon()
	case "minecraft_weather":
		return gofakeit.MinecraftWeather()
	case "minecraft_wood":
		return gofakeit.MinecraftWood()
	case "slogan":
		return gofakeit.Slogan()
	case "blurb":
		return gofakeit.Blurb()
	case "comment":
		return gofakeit.Comment()
	case "question":
		return gofakeit.Question()
	case "interjection":
		return gofakeit.Interjection()
	case "connective":
		return gofakeit.Connective()
	case "buzzword":
		return gofakeit.BuzzWord()
	case "hipster_word":
		return gofakeit.HipsterWord()
	case "hipster_sentence":
		return gofakeit.HipsterSentence(5)
	case "hipster_paragraph":
		return gofakeit.HipsterParagraph(5, 10, 3, "\n")
	case "hacker_phrase":
		return gofakeit.HackerPhrase()
	case "hacker_abbreviation":
		return gofakeit.HackerAbbreviation()
	case "hacker_adjective":
		return gofakeit.HackerAdjective()
	case "hacker_noun":
		return gofakeit.HackerNoun()
	case "hacker_verb":
		return gofakeit.HackerVerb()
	case "hackering_verb":
		return gofakeit.HackeringVerb()
	case "lorem_ipsum_word":
		return gofakeit.LoremIpsumWord()
	case "lorem_ipsum_sentence":
		return gofakeit.LoremIpsumSentence(5)
	case "lorem_ipsum_paragraph":
		return gofakeit.LoremIpsumParagraph(5, 10, 3, "\n")
	case "flip_coin":
		return gofakeit.FlipACoin()
	case "dice":
		return gofakeit.Dice(1, []uint{1, 2, 3, 4, 5, 6})
	case "weight":
		weighted, _ := gofakeit.Weighted([]interface{}{"light", "medium", "heavy"}, []float32{0.1, 0.3, 0.6})
		return weighted
	// Units
	case "unit":
		return gofakeit.Unit()
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
	if page < 1 || perPage < 1 {
		return make([]interface{}, 0)
	}

	dataList := make([]interface{}, perPage)
	for i := 0; i < perPage; i++ {
		dataList[i] = generateData(fields)
	}

	return dataList
}
