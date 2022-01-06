package main

import (
	"fmt"
	"github.com/iamtakdir/sql-migrate-cli/cmd"
)

func main() {
	fmt.Println("Migration CLI")

	cmd.Execute()
}
