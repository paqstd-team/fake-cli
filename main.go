package main

import (
	"encoding/json"
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pagination
const PER_PAGE int = 10

func Field(field string) string {
	switch field {
	case "name":
		return gofakeit.Name()
	case "phrase":
		return gofakeit.Phrase()
	}

	return "incorrected type"
}

func Generator(endpoint map[string]interface{}) any {
	scheme := endpoint["scheme"].(map[string]interface{})

	if endpoint["response"].(string) == "single" {
		single := make(map[string]interface{})
		for key, value := range scheme {
			single[key] = Field(value.(string))
		}

		return single
	} else {
		var per_page int = PER_PAGE

		if count := endpoint["count"]; count != nil {
			per_page = int(count.(float64))
		}

		generated := make([]map[string]interface{}, per_page)
		for i := 0; i < per_page; i++ {
			single := make(map[string]interface{})
			for key, value := range scheme {
				single[key] = Field(value.(string))
			}
			generated[i] = single
		}

		return generated
	}

}

func main() {
	// set flags
	var ConfigFile string
	var Port int

	var rootCmd = &cobra.Command{
		Use:   "fake",
		Short: "Create fake api with auto-generated response.",
		Long:  "Create server with fake response data. Using config file fakeconfig.json",
		Run: func(cmd *cobra.Command, args []string) {
			viper.SetConfigFile(ConfigFile)
			viper.ReadInConfig()

			// create server instance
			server := fiber.New()

			server.All("/*", func(c *fiber.Ctx) error {
				// check current path
				data := viper.Get(fmt.Sprintf("endpoints.%v", c.Params("*")))
				if data == nil {
					return c.SendStatus(404)
				}

				endpoint := data.(map[string]interface{})
				resp, err := json.Marshal(Generator(endpoint))
				if err != nil {
					panic(err)
				}

				return c.SendString((string(resp)))
			})

			server.Listen(fmt.Sprintf(":%v", Port))
		},
	}

	// run command with flags
	rootCmd.Flags().StringVarP(&ConfigFile, "config", "c", "fake.json", "Config file for create server with schema & endpoints.")
	rootCmd.Flags().IntVarP(&Port, "port", "p", 8765, "Running port.")
	rootCmd.Execute()
}
