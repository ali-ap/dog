# dog
simple domain information groper

## Installation

You can easily run the application using docker-compose.


For building the images.
```bash
docker compose build --no-cache   
```
For Running The application.
```bash
docker-compose up --force-recreate
```

## Usage

By default, OpenAPI documentation can be accessed using the below URL.

```python
http://localhost:8083/docs
```

By default, the application frontend can be accessed using the below URL.

```python
http://localhost:8082/
```

If any of those ports (8082,8083) has been already occupied you can change them in the docker-compose file. Just make sure if you change the backend service which uses 8083, change the REACT_APP_API_URL environment variable as well.

```python
REACT_APP_API_URL: http://localhost:8083/
```

use the left-hand side drawer menu to navigate to the dig page and you can check the raw view option to see the response raw result.

![a glimpse of the ui and dns lookup page](/assets/img1.png "Dig component")

## Contributing
Pull requests are welcome. feel free to add comments and improvement request and I will change them ASAP.


## License
[MIT](https://choosealicense.com/licenses/mit/)
