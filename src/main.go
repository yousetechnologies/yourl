package main

import (
	"yourl/server"
)

func main() {
	srv := new(server.HttpServer);
	srv.Run();
}
