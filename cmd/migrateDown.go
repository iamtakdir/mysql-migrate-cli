package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iamtakdir/sql-migrate-cli/database"
	"github.com/spf13/cobra"
)


var migrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate Down command installs v2 -> v1",
	Long:  "migrate Down command handles v2 -> v1 of database ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running migration Down")
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
			fmt.Println("Downgrading Error, Database is not working", err)
		}

		if err = migratenow.Down(); err!=nil{
			fmt.Println("Migrate Down ERROR ", err)
		}

	},
}

func init() {
	migrateCmd.AddCommand(migrateDownCmd)
}
