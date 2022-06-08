# 📆 CalendarWebServer

###  💻You can find the client on this [repository](https://github.com/PiterWeb/CalendarWeb)  💻

## Purpouse

Nowadays it is dificult to save your calendar in one place and still have privacy 🔐, so I tried to make a full-stack Calendar 🌐 server + 💻 client (web) which anyone can host to use it.

## Technologies used 📘

### Golang (Go)

- Fiber 🔗 (http server - inspired on expressjs)
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

    docker build --tag calendar-server .

    docker run -p 8080 calendar-server
