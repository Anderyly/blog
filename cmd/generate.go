package main

import "blog/cmd/sql"

func main() {
	if err := sql.Cmd().Execute(); err != nil {
		panic(err)
	}
}
