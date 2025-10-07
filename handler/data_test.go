package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paqstd-team/fake-cli/v2/config"
)

func TestData_GenerateFields(t *testing.T) {
	tests := []struct {
		name         string
		fields       map[string]any
		expectedKeys []string
	}{
		{
			name: "user_profile_fields",
			fields: map[string]any{
				"id":         "uuid",
				"username":   "username",
				"email":      "email",
				"first_name": "first_name",
				"last_name":  "last_name",
				"phone":      "phone",
				"created_at": "date",
			},
			expectedKeys: []string{"id", "username", "email", "first_name", "last_name", "phone", "created_at"},
		},
		{
			name: "product_catalog_fields",
			fields: map[string]any{
				"product_id":  "uuid",
				"title":       "word",
				"description": "sentence",
				"price":       "float",
				"category":    "word",
				"in_stock":    "bool",
				"tags":        []any{"word", "word", "word"},
			},
			expectedKeys: []string{"product_id", "title", "description", "price", "category", "in_stock", "tags"},
		},
		{
			name: "geographic_data_fields",
			fields: map[string]any{
				"location_id": "uuid",
				"city":        "city",
				"state":       "state",
				"country":     "country",
				"latitude":    "latitude",
				"longitude":   "longitude",
				"postal_code": "zip",
			},
			expectedKeys: []string{"location_id", "city", "state", "country", "latitude", "longitude", "postal_code"},
		},
		{
			name: "nested_structure_fields",
			fields: map[string]any{
				"order_id": "uuid",
				"customer": map[string]any{
					"id":    "uuid",
					"name":  "name",
					"email": "email",
				},
				"items": []any{
					map[string]any{
						"product_id": "uuid",
						"quantity":   "int",
						"price":      "float",
					},
				},
				"total":  "float",
				"status": "word",
			},
			expectedKeys: []string{"order_id", "customer", "items", "total", "status"},
		},
		{
			name: "mixed_data_types",
			fields: map[string]any{
				"id":            "uuid",
				"text_content":  "paragraph",
				"short_text":    "word",
				"numeric_value": "int",
				"decimal_value": "float",
				"boolean_flag":  "bool",
				"url_link":      "url",
				"ip_address":    "ip",
				"domain_name":   "domain",
			},
			expectedKeys: []string{"id", "text_content", "short_text", "numeric_value", "decimal_value", "boolean_flag", "url_link", "ip_address", "domain_name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL:      "/api/data",
						Response: tt.fields,
					},
				},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(http.MethodGet, "/api/data", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var response map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for _, key := range tt.expectedKeys {
				if _, exists := response[key]; !exists {
					t.Errorf("Expected key '%s' not found in response", key)
				}
			}

			if len(response) != len(tt.expectedKeys) {
				t.Errorf("Expected %d keys, got %d", len(tt.expectedKeys), len(response))
			}
		})
	}
}

func TestData_MapStringStringPath(t *testing.T) {
	tests := []struct {
		name         string
		fields       map[string]string
		expectedKeys []string
	}{
		{
			name: "simple_string_map",
			fields: map[string]string{
				"id":    "uuid",
				"name":  "name",
				"email": "email",
			},
			expectedKeys: []string{"id", "name", "email"},
		},
		{
			name: "location_string_map",
			fields: map[string]string{
				"city":    "city",
				"state":   "state",
				"country": "country",
			},
			expectedKeys: []string{"city", "state", "country"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				Endpoints: []config.Endpoint{
					{
						URL:      "/api/simple",
						Response: tt.fields,
					},
				},
			}
			handler := MakeHandler(cfg)

			req := httptest.NewRequest(http.MethodGet, "/api/simple", nil)
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", w.Code)
			}

			var response map[string]any
			if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
				t.Fatalf("Failed to unmarshal response: %v", err)
			}

			for _, key := range tt.expectedKeys {
				if _, exists := response[key]; !exists {
					t.Errorf("Expected key '%s' not found in response", key)
				}
			}
		})
	}
}

func TestData_AllFieldTypes(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/all-fields",
				Response: map[string]any{
					// Identifiers
					"uuid": "uuid",
					"ssn":  "ssn",

					// Geographic
					"city":          "city",
					"state":         "state",
					"country":       "country",
					"latitude":      "latitude",
					"longitude":     "longitude",
					"address":       "address",
					"street":        "street",
					"zip":           "zip",
					"postal_code":   "postal_code",
					"timezone":      "timezone",
					"timezone_abbr": "timezone_abbr",
					"timezone_full": "timezone_full",

					// Personal
					"name":        "name",
					"name_prefix": "name_prefix",
					"name_suffix": "name_suffix",
					"first_name":  "first_name",
					"last_name":   "last_name",
					"gender":      "gender",
					"hobby":       "hobby",

					// Contact
					"email": "email",
					"phone": "phone",

					// Authentication
					"username": "username",
					"password": "password",

					// Text Content
					"paragraph":             "paragraph",
					"sentence":              "sentence",
					"phrase":                "phrase",
					"quote":                 "quote",
					"word":                  "word",
					"blurb":                 "blurb",
					"comment":               "comment",
					"question":              "question",
					"interjection":          "interjection",
					"connective":            "connective",
					"buzzword":              "buzzword",
					"hipster_word":          "hipster_word",
					"hipster_sentence":      "hipster_sentence",
					"hipster_paragraph":     "hipster_paragraph",
					"lorem_ipsum_word":      "lorem_ipsum_word",
					"lorem_ipsum_sentence":  "lorem_ipsum_sentence",
					"lorem_ipsum_paragraph": "lorem_ipsum_paragraph",

					// Time & Date
					"date":         "date",
					"datetime":     "datetime",
					"time":         "time",
					"second":       "second",
					"minute":       "minute",
					"hour":         "hour",
					"month":        "month",
					"day":          "day",
					"year":         "year",
					"weekday":      "weekday",
					"month_string": "month_string",

					// Web & Network
					"url":                "url",
					"domain":             "domain",
					"domain_name":        "domain_name",
					"domain_suffix":      "domain_suffix",
					"ip":                 "ip",
					"ipv4":               "ipv4",
					"ipv6":               "ipv6",
					"mac_address":        "mac_address",
					"http_method":        "http_method",
					"http_status_code":   "http_status_code",
					"http_status_simple": "http_status_simple",
					"user_agent":         "user_agent",
					"chrome_user_agent":  "chrome_user_agent",
					"firefox_user_agent": "firefox_user_agent",
					"safari_user_agent":  "safari_user_agent",
					"opera_user_agent":   "opera_user_agent",

					// Numbers
					"int":           "int",
					"int8":          "int8",
					"int16":         "int16",
					"int32":         "int32",
					"int64":         "int64",
					"uint8":         "uint8",
					"uint16":        "uint16",
					"uint32":        "uint32",
					"uint64":        "uint64",
					"float":         "float",
					"float32":       "float32",
					"float64":       "float64",
					"bool":          "bool",
					"number":        "number",
					"int_n":         "int_n",
					"uint_n":        "uint_n",
					"float32_range": "float32_range",
					"float64_range": "float64_range",
					"digit":         "digit",
					"digit_n":       "digit_n",
					"letter":        "letter",
					"letter_n":      "letter_n",
					"vowel":         "vowel",

					// Financial
					"price":            "price",
					"currency":         "currency",
					"currency_long":    "currency_long",
					"currency_code":    "currency_code",
					"credit_card":      "credit_card",
					"credit_card_cvv":  "credit_card_cvv",
					"credit_card_exp":  "credit_card_exp",
					"credit_card_type": "credit_card_type",
					"cvv":              "cvv",
					"cvc":              "cvc",
					"expiry":           "expiry",
					"expiration":       "expiration",

					// Banking
					"bank_name":   "bank_name",
					"bank_type":   "bank_type",
					"ein":         "ein",
					"ach_account": "ach_account",
					"ach_routing": "ach_routing",

					// Business
					"company":        "company",
					"job_title":      "job_title",
					"job_descriptor": "job_descriptor",
					"job_level":      "job_level",
					"bs":             "bs",

					// Products
					"product_name":        "product_name",
					"product_category":    "product_category",
					"product_description": "product_description",
					"product_feature":     "product_feature",
					"product_material":    "product_material",
					"product_upc":         "product_upc",
					"product_audience":    "product_audience",
					"product_benefit":     "product_benefit",
					"product_dimension":   "product_dimension",
					"product_isbn":        "product_isbn",
					"product_suffix":      "product_suffix",
					"product_use_case":    "product_use_case",
					"brand":               "brand",
					"color":               "color",
					"hex_color":           "hex_color",
					"rgb_color":           "rgb_color",
					"safe_color":          "safe_color",

					// Animals
					"animal":      "animal",
					"animal_type": "animal_type",
					"bird":        "bird",
					"cat":         "cat",
					"dog":         "dog",
					"farm_animal": "farm_animal",
					"pet_name":    "pet_name",

					// Food
					"breakfast": "breakfast",
					"lunch":     "lunch",
					"dinner":    "dinner",
					"snack":     "snack",
					"dessert":   "dessert",
					"drink":     "drink",
					"fruit":     "fruit",
					"vegetable": "vegetable",

					// Beer
					"beer_name":    "beer_name",
					"beer_style":   "beer_style",
					"beer_hop":     "beer_hop",
					"beer_malt":    "beer_malt",
					"beer_yeast":   "beer_yeast",
					"beer_alcohol": "beer_alcohol",
					"beer_blg":     "beer_blg",
					"beer_ibu":     "beer_ibu",

					// Cars
					"car_maker":             "car_maker",
					"car_model":             "car_model",
					"car_type":              "car_type",
					"car_fuel_type":         "car_fuel_type",
					"car_transmission_type": "car_transmission_type",

					// Entertainment
					"movie_name":  "movie_name",
					"movie_genre": "movie_genre",
					"book_title":  "book_title",
					"book_author": "book_author",
					"book_genre":  "book_genre",

					// Music
					"song":        "song",
					"song_artist": "song_artist",
					"song_genre":  "song_genre",
					"song_name":   "song_name",

					// Technology
					"app_name":    "app_name",
					"app_author":  "app_author",
					"app_version": "app_version",

					// Cryptocurrency
					"bitcoin_address":     "bitcoin_address",
					"bitcoin_private_key": "bitcoin_private_key",

					// Gaming
					"gamertag": "gamertag",

					// Minecraft
					"minecraft_animal":           "minecraft_animal",
					"minecraft_armor_part":       "minecraft_armor_part",
					"minecraft_armor_tier":       "minecraft_armor_tier",
					"minecraft_biome":            "minecraft_biome",
					"minecraft_dye":              "minecraft_dye",
					"minecraft_food":             "minecraft_food",
					"minecraft_mob_boss":         "minecraft_mob_boss",
					"minecraft_mob_hostile":      "minecraft_mob_hostile",
					"minecraft_mob_neutral":      "minecraft_mob_neutral",
					"minecraft_mob_passive":      "minecraft_mob_passive",
					"minecraft_ore":              "minecraft_ore",
					"minecraft_tool":             "minecraft_tool",
					"minecraft_villager_job":     "minecraft_villager_job",
					"minecraft_villager_level":   "minecraft_villager_level",
					"minecraft_villager_station": "minecraft_villager_station",
					"minecraft_weapon":           "minecraft_weapon",
					"minecraft_weather":          "minecraft_weather",
					"minecraft_wood":             "minecraft_wood",

					// Miscellaneous
					"emoji":             "emoji",
					"emoji_alias":       "emoji_alias",
					"emoji_category":    "emoji_category",
					"emoji_description": "emoji_description",
					"emoji_tag":         "emoji_tag",
					"slogan":            "slogan",
					"flip_coin":         "flip_coin",
					"dice":              "dice",
					"weight":            "weight",

					// Units
					"unit": "unit",

					// Unknown field test
					"unknown_field": "unknown_value",
				},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/all-fields", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedFields := []string{
		// Identifiers
		"uuid", "ssn",

		// Geographic
		"city", "state", "country", "latitude", "longitude", "address", "street", "zip", "postal_code",
		"timezone", "timezone_abbr", "timezone_full",

		// Personal
		"name", "name_prefix", "name_suffix", "first_name", "last_name", "gender", "hobby",

		// Contact
		"email", "phone",

		// Authentication
		"username", "password",

		// Text Content
		"paragraph", "sentence", "phrase", "quote", "word", "blurb", "comment", "question",
		"interjection", "connective", "buzzword", "hipster_word", "hipster_sentence", "hipster_paragraph",
		"lorem_ipsum_word", "lorem_ipsum_sentence", "lorem_ipsum_paragraph",

		// Time & Date
		"date", "datetime", "time", "second", "minute", "hour", "month", "day", "year", "weekday", "month_string",

		// Web & Network
		"url", "domain", "domain_name", "domain_suffix", "ip", "ipv4", "ipv6", "mac_address",
		"http_method", "http_status_code", "http_status_simple", "user_agent", "chrome_user_agent",
		"firefox_user_agent", "safari_user_agent", "opera_user_agent",

		// Numbers
		"int", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64",
		"float", "float32", "float64", "bool", "number", "int_n", "uint_n", "float32_range",
		"float64_range", "digit", "digit_n", "letter", "letter_n", "vowel",

		// Financial
		"price", "currency", "currency_long", "currency_code", "credit_card", "credit_card_cvv",
		"credit_card_exp", "credit_card_type", "cvv", "cvc", "expiry", "expiration",

		// Banking
		"bank_name", "bank_type", "ein", "ach_account", "ach_routing",

		// Business
		"company", "job_title", "job_descriptor", "job_level", "bs",

		// Products
		"product_name", "product_category", "product_description", "product_feature", "product_material",
		"product_upc", "product_audience", "product_benefit", "product_dimension", "product_isbn",
		"product_suffix", "product_use_case", "brand", "color", "hex_color", "rgb_color", "safe_color",

		// Animals
		"animal", "animal_type", "bird", "cat", "dog", "farm_animal", "pet_name",

		// Food
		"breakfast", "lunch", "dinner", "snack", "dessert", "drink", "fruit", "vegetable",

		// Beer
		"beer_name", "beer_style", "beer_hop", "beer_malt", "beer_yeast", "beer_alcohol", "beer_blg", "beer_ibu",

		// Cars
		"car_maker", "car_model", "car_type", "car_fuel_type", "car_transmission_type",

		// Entertainment
		"movie_name", "movie_genre", "book_title", "book_author", "book_genre",

		// Music
		"song", "song_artist", "song_genre", "song_name",

		// Technology
		"app_name", "app_author", "app_version",

		// Cryptocurrency
		"bitcoin_address", "bitcoin_private_key",

		// Gaming
		"gamertag",

		// Minecraft
		"minecraft_animal", "minecraft_armor_part", "minecraft_armor_tier", "minecraft_biome",
		"minecraft_dye", "minecraft_food", "minecraft_mob_boss", "minecraft_mob_hostile",
		"minecraft_mob_neutral", "minecraft_mob_passive", "minecraft_ore", "minecraft_tool",
		"minecraft_villager_job", "minecraft_villager_level", "minecraft_villager_station",
		"minecraft_weapon", "minecraft_weather", "minecraft_wood",

		// Miscellaneous
		"emoji", "emoji_alias", "emoji_category", "emoji_description", "emoji_tag", "slogan",
		"flip_coin", "dice", "weight",

		// Units
		"unit",

		// Unknown field test
		"unknown_field",
	}

	for _, field := range expectedFields {
		if _, exists := response[field]; !exists {
			t.Errorf("Expected field '%s' not found in response", field)
		}
	}

	if response["unknown_field"] != "unknown_value" {
		t.Errorf("Expected unknown_field to be 'unknown_value', got %v", response["unknown_field"])
	}
}

func TestData_ComplexNestedStructures(t *testing.T) {
	cfg := config.Config{
		Endpoints: []config.Endpoint{
			{
				URL: "/api/complex",
				Response: map[string]any{
					"company": map[string]any{
						"id":   "uuid",
						"name": "company",
						"address": map[string]any{
							"street": "street",
							"city":   "city",
							"zip":    "zip",
						},
						"employees": []any{
							map[string]any{
								"id":       "uuid",
								"name":     "name",
								"position": "job_title",
							},
						},
					},
					"metadata": map[string]any{
						"created_at": "date",
						"version":    "int",
						"active":     "bool",
					},
				},
			},
		},
	}
	handler := MakeHandler(cfg)

	req := httptest.NewRequest(http.MethodGet, "/api/complex", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", w.Code)
	}

	var response map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if _, exists := response["company"]; !exists {
		t.Error("Expected 'company' key not found")
	}

	if _, exists := response["metadata"]; !exists {
		t.Error("Expected 'metadata' key not found")
	}

	company, ok := response["company"].(map[string]any)
	if !ok {
		t.Fatal("Expected 'company' to be a map")
	}

	if _, exists := company["address"]; !exists {
		t.Error("Expected 'address' key not found in company")
	}

	if _, exists := company["employees"]; !exists {
		t.Error("Expected 'employees' key not found in company")
	}
}

func TestData_GenerateFieldCoverage(t *testing.T) {
	// Test all possible field types to ensure 100% coverage
	testCases := []string{
		// Geography
		"address", "street", "zip", "postal_code", "timezone", "timezone_abbr", "timezone_full",

		// Time
		"datetime", "time", "weekday", "month_string",

		// Financial
		"price", "currency", "currency_long", "currency_code", "credit_card", "credit_card_cvv",
		"credit_card_exp", "credit_card_type", "cvv", "cvc", "expiry", "expiration",

		// Banking
		"bank_name", "bank_type", "ein", "ach_account", "ach_routing",

		// Internet
		"domain_name", "domain_suffix", "ipv4", "ipv6", "mac_address", "http_method",
		"http_status_code", "http_status_simple", "user_agent", "chrome_user_agent",
		"firefox_user_agent", "safari_user_agent", "opera_user_agent",

		// Products
		"product_name", "product_category", "product_description", "product_feature", "product_material",
		"product_upc", "product_audience", "product_benefit", "product_dimension", "product_isbn",
		"product_suffix", "product_use_case", "brand", "color", "hex_color", "rgb_color", "safe_color",

		// Animals
		"animal", "animal_type", "bird", "cat", "dog", "farm_animal", "pet_name",

		// Food
		"breakfast", "lunch", "dinner", "snack", "dessert", "drink", "fruit", "vegetable",

		// Beer
		"beer_name", "beer_style", "beer_hop", "beer_malt", "beer_yeast", "beer_alcohol", "beer_blg", "beer_ibu",

		// Cars
		"car_maker", "car_model", "car_type", "car_fuel_type", "car_transmission_type",

		// Entertainment
		"movie_name", "movie_genre", "book_title", "book_author", "book_genre",

		// Music
		"song", "song_artist", "song_genre", "song_name",

		// Apps
		"app_name", "app_author", "app_version",

		// Numbers
		"int", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64",
		"float", "float32", "float64", "bool", "number", "int_n", "uint_n", "float32_range",
		"float64_range", "digit", "digit_n", "letter", "letter_n", "vowel",

		// Crypto
		"bitcoin_address", "bitcoin_private_key",

		// Gaming
		"gamertag",

		// Minecraft
		"minecraft_animal", "minecraft_armor_part", "minecraft_armor_tier", "minecraft_biome",
		"minecraft_dye", "minecraft_food", "minecraft_mob_boss", "minecraft_mob_hostile",
		"minecraft_mob_neutral", "minecraft_mob_passive", "minecraft_ore", "minecraft_tool",
		"minecraft_villager_job", "minecraft_villager_level", "minecraft_villager_station",
		"minecraft_weapon", "minecraft_weather", "minecraft_wood",

		// Miscellaneous
		"emoji", "emoji_alias", "emoji_category", "emoji_description", "emoji_tag", "slogan",
		"blurb", "comment", "question", "interjection", "connective", "buzzword", "hipster_word",
		"hipster_sentence", "hipster_paragraph", "hacker_phrase", "hacker_abbreviation", "hacker_adjective",
		"hacker_noun", "hacker_verb", "hackering_verb", "lorem_ipsum_word", "lorem_ipsum_sentence",
		"lorem_ipsum_paragraph", "flip_coin", "dice", "weight",

		// Units
		"unit",

		// Business
		"company", "job_title", "job_descriptor", "job_level", "bs",
	}

	for _, fieldType := range testCases {
		t.Run(fieldType, func(t *testing.T) {
			result := generateField(fieldType)
			if result == nil {
				t.Errorf("generateField(%s) returned nil", fieldType)
			}
		})
	}
}
