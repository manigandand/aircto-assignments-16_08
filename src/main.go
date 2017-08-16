package main

import "os"

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"))

	a.Run(":3011")
}
