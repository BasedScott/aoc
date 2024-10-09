package year2023

import "os"

func Env() {
	os.Setenv("YEAR2023", os.Getenv("GOPATH"))
}
