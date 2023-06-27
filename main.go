package main

import (
	"github.com/yu5429/mycoin/cli"
	"github.com/yu5429/mycoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
