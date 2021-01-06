package main

func main() {
	app := &app{}
	app.init()
	app.run(":8080")
}
