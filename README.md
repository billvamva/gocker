# gocker
a simple, fully tested, poker game implementation using Go, creating a web server and a cli.

The project includes a very simple frontend that interacts with the backend using a Websocket.

docker build -t gocker .

docker run -p 5000:5000 gocker
