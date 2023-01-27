# Fake-Cli

## About
Generate fake-data and create local web-server. 

Cli options:  
- fake  
Run cli with default options.  

- fake -p 4321  
Run cli with custom port `localhost:4321`. By default listen `localhost:8765`.

- fake -c custom.json  
Run cli with custom file configuration. By default config file is `fake.json`. 


## Config file
### Example fake.json file
```
{
    "endpoints": {
        "users/": {
            "response": "some",
            "scheme": {
                "name": "name",
                "surname": "phrase"
            }
        },
        "users/1/": {
            "response": "single",
            "scheme": {
                "name": "name"
            }
        }
    }
}
```

- endpoints:  
An object of all endpoints to run the web server.

- - key:  
Each key is a path in the browser where one or another endpoint can be requested

- - - response:  
Points to the return type. It can be "single" or "some".

- - - count:  
If response == "some" is specified. Specifies how many records to return.

- - - scheme:  
The schema of the returned data. Fake data will be generated for the specified fields. Each key is a field name, each value is a generated type.


### Available schema type values
- name  
Generation of first and last name. Example: `Markus Moen`.

- phrase  
Random phrase generation. Example: `head in the sand`.
