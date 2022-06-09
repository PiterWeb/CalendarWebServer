# 📆 CalendarWebServer

###  💻You can find the client on this [repository](https://github.com/PiterWeb/CalendarWeb)  💻

## Purpouse

Nowadays it is dificult to save your calendar in one place and still have privacy 🔐, so I tried to make a full-stack Calendar 🌐 server + 💻 client (web) which anyone can host to use it.

## Before runing the project

Set the next env variables on a .env file located at the root folder of the project

    SECRET_CRYPT = <secret>
    SECRET_COOKIE = <secret>
    PORT = <port>
    MONGO_USER =  <mongodb user>
    MONGO_PSW = <mongdb password>
    MONGO_DM = <mongodb domain>

## Technologies used 📘

### Golang (Go)

- Fiber 🔗 (backend framework - inspired on expressjs)
	#### Middlewares
	- Compress  (gzip response)
	- Encrypt Cookie
	- Cors (modify CORS)
	- Etag (efficient cache)
- Godotenv 🔒 (use enviroment variables)
- Mongo Driver 📦 (MongoDB driver)
- ShortID 🔡 (generate unique IDs)

	#### Build

	    go build main.go

	    ./<executable>

### Docker

- Go image : 1.17-alpine

	#### Build

	    docker build --tag calendar-server .

	    docker run -p 8080 calendar-server
