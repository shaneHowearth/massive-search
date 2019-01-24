// main.go

package main

func main() {
	a := RESTApp{}
	a.Initialise()
	// REST server will listen on port 8080
	a.Run(":8000")
}
