package lister

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var interactive = terminal.IsTerminal(int(os.Stdout.Fd()))

// ListFolders lists folders in a Grafana instance
func ListFolders() error {
	uids, err := listFolders()
	if err != nil {
		return err
	}
	for _, uid := range uids {
		fmt.Println(uid)
	}
	return nil
}

// ListDashboards lists dashboard IDs in a Grafana instance
func ListDashboards(folderID int) error {
	uids, err := listDashboards(folderID)
	if err != nil {
		return err
	}
	for _, uid := range uids {
		fmt.Println(uid)
	}
	return nil
}

// ListDatasources lists datasource IDs in a Grafana instance
func ListDatasources() error {
	uids, err := listDatasources()
	if err != nil {
		return err
	}
	for _, uid := range uids {
		fmt.Println(uid)
	}
	return nil
}
