# Fake API (CLI)

This is a command-line tool for generating fake API responses based on a JSON configuration file. The tool is written in Go and uses the [gofakeit v7](https://github.com/brianvoe/gofakeit) library to generate random data.

## Installation

To install the fakeapi command-line tool, you must have Go installed on your system. You can download and install Go from the official Go website.

Once you have Go installed, you can install the fakeapi tool by running the following command:
```bash
go install -v github.com/paqstd-team/fake-cli@latest
```

Also you can install the binary by downloading it from [one of the latest releases](https://github.com/paqstd-team/fake-cli/releases).

To generate fake API responses using this configuration file, you can run the fake command:

```bash
fake-cli -c path/to/config.json
```
Replace path/to/config.json with the path to your configuration file.

By default, the fake command starts a web server on port 8000 that responds to requests for the endpoints defined in your configuration file. You can specify a different port by adding the --p flag followed by the desired port number:

```bash
fake-cli -c path/to/config.json -p 8080
```

## Usage

To generate fake API responses, you must create a configuration file in JSON format that defines the endpoints and response template for each endpoint. We recommended use `.json` type for files. Here's an example configuration file `config.json`:

```json
{
  "endpoints": [
    {
      "url": "/users",
      "type": "GET",
      "cache": 5,
      "response": [
        {
          "id": "uuid",
          "name": "name",
          "email": "email"
        }
      ]
    },
    {
      "url": "/products",
      "cache": 10,
      "response": {
        "id": "uuid",
        "name": "word",
        "price": "float"
      }
    },
    {
      "url": "/products/{id}",
      "cache": 3,
      "response": {
        "id": "uuid",
        "tags": [{
          "id": "uuid",
          "name": "name"
        }],
        "details": {
          "about": {
            "author": "name",
            "customValue": "my-custom-value"
          }
        }
      }
    },
    {
      "url": "/list",
      "response": {
        "names": ["name", "name", "name"]
      }
    },
    {
      "url": "/submit",
      "type": "POST",
      "response": {"status": "word"}
    },
    {
      "url": "/update/{id}",
      "type": "PATCH",
      "cache": 1,
      "response": {"updated": "word"}
    },
    {
      "url": "/replace/{id}",
      "type": "PUT",
      "response": {"replaced": "word"}
    },
    {
      "url": "/remove/{id}",
      "type": "DELETE",
      "response": {"removed": "word"}
    }
  ]
}
```
**Configuration explanation:**
- `/users` - Returns a list (top-level array in `response`) with pagination (by default `page=1` and `per_page=10`), cached for 5 requests
- `/products` - Returns a single object (when `response` is an object), cached for 10 requests  
- `/products/{id}` - Individual product endpoint, cached for 3 requests
- `/list` - No caching, generates new data on every request
- `/submit` - POST endpoint, typically doesn't need caching
- `/update/{id}` - PATCH endpoint, cached for 1 request
- `/replace/{id}` - PUT endpoint, no caching
- `/remove/{id}` - DELETE endpoint, no caching

Endpoints may specify an HTTP method using `type` and support: `GET` (default), `POST`, `PATCH`, `PUT`, `DELETE`.

## Caching

Each endpoint can have its own individual cache configuration:

- **`cache: 5`** - Cache responses for 5 requests, then generate new data
- **`cache: 10`** - Cache responses for 10 requests, then generate new data  
- **`cache: 1`** - Cache only 1 response, then generate new data
- **No `cache` field** - No caching, generate new data on every request

**Important notes:**
- Caching only works for `GET` requests
- Each endpoint has its own separate cache instance
- Cache is based on request URL and query parameters
- If `cache` is not specified, the endpoint will not use caching

You can specify arrays or objects inside `response`. A top-level object means a single-object response; a top-level array (e.g., `[ { ... } ]`) means a list response where the first item defines the item template. Nested arrays/objects are supported.

## Available Data Types

| Category | Type | Description |
|----------|------|-------------|
| **Identifiers** | `uuid` | Universally unique identifier |
| | `ssn` | Social Security Number |
| **Geographic** | `city` | City name |
| | `state` | State/Province name |
| | `country` | Country name |
| | `latitude` | Geographic latitude |
| | `longitude` | Geographic longitude |
| | `address` | Full address |
| | `street` | Street name |
| | `zip` | ZIP/Postal code |
| | `postal_code` | Postal code |
| | `timezone` | Time zone |
| | `timezone_abbr` | Time zone abbreviation |
| | `timezone_full` | Full time zone name |
| **Personal** | `name` | Full name |
| | `name_prefix` | Name prefix (Mr., Mrs., Dr.) |
| | `name_suffix` | Name suffix (Jr., Sr., III) |
| | `first_name` | First name |
| | `last_name` | Last name |
| | `gender` | Gender |
| | `hobby` | Hobby/Interest |
| **Contact** | `email` | Email address |
| | `phone` | Phone number |
| **Authentication** | `username` | Username |
| | `password` | Password |
| **Text Content** | `paragraph` | Multi-sentence paragraph |
| | `sentence` | Single sentence |
| | `phrase` | Short phrase |
| | `quote` | Famous quote |
| | `word` | Single word |
| | `blurb` | Short description |
| | `comment` | Comment text |
| | `question` | Question text |
| | `interjection` | Interjection |
| | `connective` | Connective word |
| | `buzzword` | Business buzzword |
| | `hipster_word` | Hipster word |
| | `hipster_sentence` | Hipster sentence |
| | `hipster_paragraph` | Hipster paragraph |
| | `lorem_ipsum_word` | Lorem ipsum word |
| | `lorem_ipsum_sentence` | Lorem ipsum sentence |
| | `lorem_ipsum_paragraph` | Lorem ipsum paragraph |
| **Time & Date** | `date` | Date (YYYY-MM-DD) |
| | `datetime` | Date and time |
| | `time` | Time |
| | `second` | Second (0-59) |
| | `minute` | Minute (0-59) |
| | `hour` | Hour (0-23) |
| | `month` | Month (1-12) |
| | `day` | Day of month (1-31) |
| | `year` | Year |
| | `weekday` | Day of week |
| | `month_string` | Month name |
| **Web & Network** | `url` | URL |
| | `domain` | Domain name |
| | `domain_name` | Domain name only |
| | `domain_suffix` | Domain suffix (.com, .org) |
| | `ip` | IPv4 address |
| | `ipv4` | IPv4 address |
| | `ipv6` | IPv6 address |
| | `mac_address` | MAC address |
| | `http_method` | HTTP method |
| | `http_status_code` | HTTP status code |
| | `http_status_simple` | Simple HTTP status |
| | `user_agent` | User agent string |
| | `chrome_user_agent` | Chrome user agent |
| | `firefox_user_agent` | Firefox user agent |
| | `safari_user_agent` | Safari user agent |
| | `opera_user_agent` | Opera user agent |
| **Numbers** | `int` | Integer (32-bit) |
| | `int8` | 8-bit integer |
| | `int16` | 16-bit integer |
| | `int32` | 32-bit integer |
| | `int64` | 64-bit integer |
| | `uint8` | 8-bit unsigned integer |
| | `uint16` | 16-bit unsigned integer |
| | `uint32` | 32-bit unsigned integer |
| | `uint64` | 64-bit unsigned integer |
| | `float` | Float (32-bit) |
| | `float32` | 32-bit float |
| | `float64` | 64-bit float |
| | `bool` | Boolean value |
| | `number` | Random number (1-100) |
| | `digit` | Single digit |
| | `letter` | Single letter |
| | `vowel` | Vowel letter |
| | `int_n` | Random integer with N digits |
| | `uint_n` | Random unsigned integer with N digits |
| | `float32_range` | Random float32 in range |
| | `float64_range` | Random float64 in range |
| | `digit_n` | N random digits |
| | `letter_n` | N random letters |
| **Financial** | `price` | Price value |
| | `currency` | Currency code |
| | `currency_long` | Full currency name |
| | `currency_code` | Currency code |
| | `credit_card` | Credit card number |
| | `credit_card_cvv` | Credit card CVV |
| | `credit_card_exp` | Credit card expiry |
| | `credit_card_type` | Credit card type |
| | `cvv` | CVV code |
| | `cvc` | CVC code |
| | `expiry` | Expiry date |
| | `expiration` | Expiration date |
| **Banking** | `bank_name` | Bank name |
| | `bank_type` | Bank type |
| | `ein` | Employer Identification Number |
| | `ach_account` | ACH account number |
| | `ach_routing` | ACH routing number |
| **Business** | `company` | Company name |
| | `job_title` | Job title |
| | `job_descriptor` | Job descriptor |
| | `job_level` | Job level |
| | `bs` | Business speak |
| **Products** | `product_name` | Product name |
| | `product_category` | Product category |
| | `product_description` | Product description |
| | `product_feature` | Product feature |
| | `product_material` | Product material |
| | `product_upc` | Product UPC code |
| | `product_audience` | Product target audience |
| | `product_benefit` | Product benefit |
| | `product_dimension` | Product dimension |
| | `product_isbn` | Product ISBN |
| | `product_suffix` | Product suffix |
| | `product_use_case` | Product use case |
| | `brand` | Brand name |
| | `color` | Color name |
| | `hex_color` | Hex color code |
| | `rgb_color` | RGB color values |
| | `safe_color` | Safe color name |
| **Animals** | `animal` | Animal name |
| | `animal_type` | Animal type |
| | `bird` | Bird name |
| | `cat` | Cat breed |
| | `dog` | Dog breed |
| | `farm_animal` | Farm animal |
| | `pet_name` | Pet name |
| **Food** | `breakfast` | Breakfast food |
| | `lunch` | Lunch food |
| | `dinner` | Dinner food |
| | `snack` | Snack food |
| | `dessert` | Dessert |
| | `drink` | Drink |
| | `fruit` | Fruit |
| | `vegetable` | Vegetable |
| **Beer** | `beer_name` | Beer name |
| | `beer_style` | Beer style |
| | `beer_hop` | Beer hop |
| | `beer_malt` | Beer malt |
| | `beer_yeast` | Beer yeast |
| | `beer_alcohol` | Alcohol content |
| | `beer_blg` | Beer BLG |
| | `beer_ibu` | Beer IBU |
| **Cars** | `car_maker` | Car manufacturer |
| | `car_model` | Car model |
| | `car_type` | Car type |
| | `car_fuel_type` | Fuel type |
| | `car_transmission_type` | Transmission type |
| **Entertainment** | `movie_name` | Movie title |
| | `movie_genre` | Movie genre |
| | `book_title` | Book title |
| | `book_author` | Book author |
| | `book_genre` | Book genre |
| **Music** | `song` | Song name |
| | `song_artist` | Song artist |
| | `song_genre` | Song genre |
| | `song_name` | Song name |
| **Technology** | `app_name` | App name |
| | `app_author` | App author |
| | `app_version` | App version |
| **Cryptocurrency** | `bitcoin_address` | Bitcoin address |
| | `bitcoin_private_key` | Bitcoin private key |
| **Gaming** | `gamertag` | Gaming username |
| **Minecraft** | `minecraft_animal` | Minecraft animal |
| | `minecraft_armor_part` | Minecraft armor part |
| | `minecraft_armor_tier` | Minecraft armor tier |
| | `minecraft_biome` | Minecraft biome |
| | `minecraft_dye` | Minecraft dye |
| | `minecraft_food` | Minecraft food |
| | `minecraft_mob_boss` | Minecraft boss mob |
| | `minecraft_mob_hostile` | Minecraft hostile mob |
| | `minecraft_mob_neutral` | Minecraft neutral mob |
| | `minecraft_mob_passive` | Minecraft passive mob |
| | `minecraft_ore` | Minecraft ore |
| | `minecraft_tool` | Minecraft tool |
| | `minecraft_villager_job` | Minecraft villager job |
| | `minecraft_villager_level` | Minecraft villager level |
| | `minecraft_villager_station` | Minecraft villager station |
| | `minecraft_weapon` | Minecraft weapon |
| | `minecraft_weather` | Minecraft weather |
| | `minecraft_wood` | Minecraft wood type |
| **Miscellaneous** | `emoji` | Emoji |
| | `emoji_alias` | Emoji alias |
| | `emoji_category` | Emoji category |
| | `emoji_description` | Emoji description |
| | `emoji_tag` | Emoji tag |
| | `slogan` | Slogan |
| | `flip_coin` | Coin flip result |
| | `dice` | Dice roll |
| | `weight` | Weighted random selection |
| **Units** | `unit` | Unit of measurement |

## Customization

You can customize the types of fake data generated by editing the handler/handler.go file. The MakeHandler function generates fake data based on the fields and response type defined in the configuration file.

You can also add new types of fake data by modifying the switch statement in the MakeHandler function. The gofakeit library provides many built-in functions for generating fake data, and you can use these functions to generate custom data types.

## Docker
### Local
Pull image to local:
```bash
docker pull ghcr.io/paqstd-team/fake-cli
```

Run with docker:
```bash
docker run --name fake-cli -it -v ${PWD}/config.json:/app/config.json -p 8080:8080 -e CONFIG_PATH=config.json -e PORT=8080 ghcr.io/paqstd-team/fake-cli
```

### Docker Compose
Here is an example of usage `fake-cli` with docker-compose and other containers:  
```yml
services:
  # ...other services
  fake-cli:
    # pull from github container registry
    image: ghcr.io/paqstd-team/fake-cli:latest
    environment:
      # default config path is "config.json"
      - CONFIG_PATH=config.json
      # default port is 8080
      - PORT=8080
    ports:
      # link port inside container to real world
      - 8080:8080
    volumes:
      # copy config file to container
      - ./config.json:/app/config.json
```

**Example config.json for Docker:**
```json
{
  "endpoints": [
    {
      "url": "/api/users",
      "cache": 5,
      "response": {"id": "uuid", "name": "name", "email": "email"}
    },
    {
      "url": "/api/products",
      "cache": 10, 
      "response": {"id": "uuid", "title": "word", "price": "float"}
    },
    {
      "url": "/api/health",
      "response": {"status": "ok", "timestamp": "date"}
    }
  ]
}
```

## Local development
Use commands:

```bash
make        # builds the binary (default command)
make build  # builds the binary
make clean  # removes the binary
make run    # runs the program
```

## Contributing

If you find a bug or would like to suggest a new feature, you can create an issue on the GitHub repository for this project. If you'd like to contribute code, you can fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
