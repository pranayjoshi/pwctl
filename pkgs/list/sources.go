package list

import (
	"fmt"
	"os"
	"pwctl/pkgs/db"
	"text/tabwriter"

	"github.com/urfave/cli/v2"
)

func ListSources(c *cli.Context) error {
	serverAddress := c.String("c")
	db := db.NewDB()
	dbs, err := db.GetDBs(serverAddress)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "\nID\tName\tConnstr")

	for i, database := range dbs {
		fmt.Fprintf(w, "%d\t%s\t%s\n", i, database.MD_Name, database.MD_Connstr)
	}

	w.Flush()
	return nil
}
