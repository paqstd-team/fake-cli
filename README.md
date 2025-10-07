# Fake API (CLI)

This is a command-line tool for generating fake API responses based on a JSON configuration file. The tool is written in Go and uses the gofakeit library to generate random data.

## Installation

To install the fakeapi command-line tool, you must have Go installed on your system. You can download and install Go from the official Go website.

Once you have Go installed, you can install the fakeapi tool by running the following command:
```bash
go install -v github.com/paqstd-team/fake-cli@latest
```

Also you can install the binary by downloading it from [one of the latest releases](https://github.com/paqstd-team/fake-cli/releases).

## Local development
Use commands:

```bash
make        # builds the binary (default command)
make build  # builds the binary
make clean  # removes the binary
make run    # runs the program
```

## Usage

To generate fake API responses, you must create a configuration file in JSON format that defines the endpoints and response template for each endpoint. We recommended use `.json` type for files. Here's an example configuration file `config.json`:

```json
{
  "endpoints": [
    {
      "url": "/users",
      "type": "GET", // default is GET if omitted
      "cache": 5, // individual cache for this endpoint (5 requests)
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
      "cache": 10, // different cache size for this endpoint
      // single object response when response is an object
      "response": {
        "id": "uuid",
        "name": "word",
        "price": "float"
      }
    },
    {
      "url": "/products/{id}",
      "cache": 3, // smaller cache for individual product
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
      // no cache field - this endpoint won't be cached
      "response": {
        "names": ["name", "name", "name"]
      }
    },
    {
      "url": "/submit",
      "type": "POST",
      // POST endpoints typically don't need caching
      "response": {"status": "word"}
    },
    {
      "url": "/update/{id}",
      "type": "PATCH",
      "cache": 1, // cache only 1 response for updates
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
This configuration file defines several endpoints. The `/users` endpoint returns a list (top-level array in `response`) with pagination (by default `page=1` and `per_page=10`). The `/products` endpoint returns a single object (when `response` is an object). Endpoints may specify an HTTP method using `type` and support: `GET` (default), `POST`, `PATCH`, `PUT`, `DELETE`.

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
| **Personal** | `name` | Full name |
| | `name_prefix` | Name prefix (Mr., Mrs., Dr.) |
| | `name_suffix` | Name suffix (Jr., Sr., III) |
| | `first_name` | First name |
| | `last_name` | Last name |
| | `gender` | Gender |
| **Contact** | `email` | Email address |
| | `phone` | Phone number |
| **Authentication** | `username` | Username |
| | `password` | Password |
| **Text Content** | `paragraph` | Multi-sentence paragraph |
| | `sentence` | Single sentence |
| | `phrase` | Short phrase |
| | `quote` | Famous quote |
| | `word` | Single word |
| **Time & Date** | `date` | Date (YYYY-MM-DD) |
| | `second` | Second (0-59) |
| | `minute` | Minute (0-59) |
| | `hour` | Hour (0-23) |
| | `month` | Month (1-12) |
| | `day` | Day of month (1-31) |
| | `year` | Year |
| **Web & Network** | `url` | URL |
| | `domain` | Domain name |
| | `ip` | IP address |
| **Numbers** | `int` | Integer |
| | `float` | Floating point number |
| **Other** | `hobby` | Hobby/Interest |

### Type Usage Examples

```json
{
  "endpoints": [
    {
      "url": "/user-profile",
      "response": {
        "id": "uuid",           // Unique identifier
        "name": "name",         // Full name
        "email": "email",       // Email address
        "phone": "phone",       // Phone number
        "location": {
          "city": "city",       // City name
          "country": "country"  // Country name
        },
        "created_at": "date",   // Creation date
        "age": "int",           // Age as integer
        "bio": "paragraph"      // Biography text
      }
    },
    {
      "url": "/product",
      "response": {
        "id": "uuid",
        "title": "word",        // Product title
        "description": "sentence", // Product description
        "price": "float",       // Price as decimal
        "tags": ["word", "word"], // Array of tags
        "website": "url"        // Product website
      }
    }
  ]
}
```

To generate fake API responses using this configuration file, you can run the fake command:

```bash
fake-cli -c path/to/config.json
```
Replace path/to/config.json with the path to your configuration file.

By default, the fake command starts a web server on port 8000 that responds to requests for the endpoints defined in your configuration file. You can specify a different port by adding the --p flag followed by the desired port number:

```bash
fake-cli -c path/to/config.json -p 8080
```

## Migration from Global Cache

If you're upgrading from a version with global cache, here's how to migrate:

**Old configuration (global cache):**
```json
{
  "cache": 5,
  "endpoints": [...]
}
```

**New configuration (individual cache):**
```json
{
  "endpoints": [
    {
      "url": "/users",
      "cache": 5,  // Move cache setting to each endpoint
      "response": {...}
    },
    {
      "url": "/products", 
      "cache": 5,  // Apply same cache to all endpoints
      "response": {...}
    }
  ]
}
```

**Benefits of individual cache:**
- Different cache sizes for different endpoints
- Better memory management
- More granular control
- No global cache conflicts

## Examples

### Individual Cache per Endpoint

Here's how different cache configurations work in practice:

```json
{
  "endpoints": [
    {
      "url": "/api/users",
      "cache": 5,
      "response": {"id": "uuid", "name": "name"}
    },
    {
      "url": "/api/products", 
      "cache": 10,
      "response": {"id": "uuid", "title": "word"}
    },
    {
      "url": "/api/orders",
      "response": {"id": "uuid", "status": "word"}
    }
  ]
}
```

**Behavior:**
- `/api/users` - Same response for 5 requests, then new data
- `/api/products` - Same response for 10 requests, then new data  
- `/api/orders` - New data on every request (no caching)

### Cache Testing

Test your cache configuration:

```bash
# Start the server
./fake-cli config.json 8080

# Test cached endpoint (should return same data)
curl http://localhost:8080/api/users
curl http://localhost:8080/api/users  # Same response

# Test non-cached endpoint (should return different data)
curl http://localhost:8080/api/orders
curl http://localhost:8080/api/orders  # Different response
```

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

## Contributing

If you find a bug or would like to suggest a new feature, you can create an issue on the GitHub repository for this project. If you'd like to contribute code, you can fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
