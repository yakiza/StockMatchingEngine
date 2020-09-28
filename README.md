# StockMatchingEngine

Foobar is a Python library for dealing with word pluralization.

## Installation

Use the package manager [go get](https://golang.org/pkg/cmd/go/internal/get/) to install gorilla,postgres .

```bash
git clone https://github.com/yakiza/StockMatchingEngine/
```

Set environmental variables 

```bash
export APP_DB_USERNAME=postgres
export APP_DB_PASSWORD= {your password}
export APP_DB_NAME=stockmatching

```
### Run the application 

```bash
go run *.go
````

You can install using [Docker](https://www.docker.com/) 

```bash
docker-compose up --build
````

## Usage
#### Add User

```bash
curl -v  127.0.0.1:8000/order/create -d '{ "userid": 0, "ticker": "AABB", "price": 44.20, "quantity": 5, "command": "SELL"}'

```

#### View all Users

```bash
curl -v  127.0.0.1:8000/order/create 

```

#### Add Order

```bash
curl -v  127.0.0.1:8000/user/create -d '{ "firstname": "Example", "lastname": "Example}'

```

#### View all Oders

```bash
curl -v  127.0.0.1:8000/order/create 

```


#### Get ticker lower buy and higher sell value

```bash
curl -v  127.0.0.1:8000/order/{ticker name goes here} 

```

#### Read my special notes 
[MY NOTES](https://github.com/yakiza/StockMatchingEngine/blob/master/.%7Elock.Documentation.odt%23)
