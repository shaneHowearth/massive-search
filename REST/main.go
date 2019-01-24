// main.go

package main

func main() {
	a := RESTApp{}
	a.Initialise()
	a.Run(":8080")
}
