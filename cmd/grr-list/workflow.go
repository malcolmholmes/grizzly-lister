package main

import (
	"strconv"

	"github.com/go-clix/cli"
	lister "github.com/malcolmholmes/grizzly-lister/pkg/lister"
)

func foldersCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "folders",
		Short: "list folders",
		Args:  cli.ArgsExact(0),
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		return lister.ListFolders()
	}
	return cmd
}

func dashboardsCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "dashboards <folder-id>",
		Short: "list dashboards IDs in a folder",
		Args:  cli.ArgsExact(1),
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		ID, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		return lister.ListDashboards(ID)
	}
	return cmd
}

func datasourcesCmd() *cli.Command {
	cmd := &cli.Command{
		Use:   "datasources",
		Short: "list datasource IDs",
		Args:  cli.ArgsExact(0),
	}
	cmd.Run = func(cmd *cli.Command, args []string) error {
		return lister.ListDatasources()
	}
	return cmd
}
