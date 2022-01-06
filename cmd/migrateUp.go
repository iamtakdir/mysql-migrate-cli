package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iamtakdir/sql-migrate-cli/database"
	"github.com/spf13/cobra"
)

const filePath string = "file://F:/projects/go/sql-migrate-cli/migrations"

var migrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up command installs v1",
	Long:  "migrate up command handles v1 of database ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migrationUp")
		db := database.Open()
		dbDriver, err:= mysql.WithInstance(db, &mysql.Config{})
		if err != nil{
			fmt.Println("Error in creating instance ", err)
		}

		migratenow, err := migrate.NewWithDatabaseInstance(
			filePath,
			"migratedb", 
			dbDriver,
		)
		
		if err!= nil {
			fmt.Println("Migrate Error, Database is not working", err)
		}

		if err = migratenow.Up(); err!=nil{
			fmt.Println("Migrate Up ERROR ", err)
		}

	},
}

func init() {
	migrateCmd.AddCommand(migrateUpCmd)
}
