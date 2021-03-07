package lister

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Folder encapsulates a dashboard folder object from the Grafana API
type Folder struct {
	ID int `json:"id"`
}

func listFolders() ([]string, error) {
	grafanaURL, err := getGrafanaURL("api/folders?limit=1000")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(grafanaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		if resp.StatusCode >= 400 {
			return nil, errors.New(resp.Status)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var folders = []Folder{}
	if err := json.Unmarshal([]byte(string(body)), &folders); err != nil {
		return nil, err
	}
	IDs := []string{}
	for _, folder := range folders {
		IDs = append(IDs, strconv.Itoa(folder.ID))
	}
	return IDs, nil
}

// Dashboard represents a dashboard as contained within a Folder
type Dashboard struct {
	UID string `json:"uid"`
}

func listDashboards(folderID int) ([]string, error) {
	grafanaURL, err := getGrafanaURL(fmt.Sprintf("/api/search?folderIds=%d&type=dash-db", folderID))
	log.Println("URL", grafanaURL)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(grafanaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		if resp.StatusCode >= 400 {
			return nil, errors.New(resp.Status)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var dashboards []Dashboard
	if err := json.Unmarshal([]byte(string(body)), &dashboards); err != nil {
		return nil, err
	}
	UIDs := []string{}
	for _, dashboard := range dashboards {
		UIDs = append(UIDs, dashboard.UID)
	}
	return UIDs, nil
}

// Datasource represents a datasource as contained within a Folder
type Datasource struct {
	UID string `json:"name"`
}

func listDatasources() ([]string, error) {
	grafanaURL, err := getGrafanaURL("api/datasources")
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(grafanaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		if resp.StatusCode >= 400 {
			return nil, errors.New(resp.Status)
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var datasources = []Datasource{}
	if err := json.Unmarshal([]byte(string(body)), &datasources); err != nil {
		return nil, err
	}
	IDs := []string{}
	for _, datasource := range datasources {
		IDs = append(IDs, datasource.UID)
	}
	return IDs, nil
}
