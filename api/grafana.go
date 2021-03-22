package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DashboardTemplate struct {
	Dashboard DashboardDefinition `json:"dashboard"`
}

type DashboardDefinition struct {
	Style         string        `json:"style"`
	Templating    interface{}   `json:"templating"`
	Links         []interface{} `json:"links"`
	GraphTooltip  int           `json:"graphTooltip"`
	AlertPanelMap interface{}   `json:"alertPanelMap"`
	Editable      bool          `json:"editable"`
	Annotations   interface{}   `json:"annotations"`
	GnetId        interface{}   `json:"gnetId"`
	Timepicker    interface{}   `json:"timepicker"`
	Title         string        `json:"title"`
	Version       int           `json:"version"`
	Time          interface{}   `json:"time"`
	Timezone      string        `json:"timezone"`
	SchemaVersion int           `json:"schemaVersion"`
	Panels        []interface{} `json:"panels"`
}

var dashboard = "/grafana/dashboards"

func (client *HgClient) Create(dashboardTemplate *DashboardTemplate) error {
	url := client.GetResourceById(dashboard, "")

	var requestBody, _ = json.Marshal(dashboardTemplate)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	defer response.Body.Close()

	apiResponse := Response{response}

	err = apiResponse.HandleApiError("error getting dashboard")

	if err != nil {
		return fmt.Errorf("error creating dashboard: %v", err)
	}

	return nil
}

func (client *HgClient) Update(dashboardTemplate *DashboardTemplate) error {
	url := client.GetResource(dashboard)

	var requestBody, _ = json.Marshal(dashboardTemplate)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	response, err := client.HttpClient.Do(req)
	defer response.Body.Close()

	apiResponse := Response{response}

	err = apiResponse.HandleApiError("error updating dashboard")

	if err != nil {
		return fmt.Errorf("error doing PUT request: %v", err)
	}

	return nil
}

func (client *HgClient) Exist(name string) bool {
	url := client.GetResourceById(dashboard, toSlug(name))

	response, err := http.Get(url)
	defer response.Body.Close()

	var value interface{}

	if err != nil {
		return false
	}

	json.NewDecoder(response.Body).Decode(&value)

	return response.StatusCode != 404
}

func (client *HgClient) Get(name string) (dash *DashboardTemplate, err error) {
	var value DashboardTemplate
	url := client.GetResourceById(dashboard, name)

	response, err := http.Get(url)
	defer response.Body.Close()

	apiResponse := Response{response}

	err = apiResponse.HandleApiError(fmt.Sprintf("error getting dashboard %s", name))

	if err != nil {
		return nil, fmt.Errorf("get dashboard %s: %v", name, err)
	}

	json.NewDecoder(response.Body).Decode(&value)

	return &value, nil
}

func (client *HgClient) Delete(id string) error {
	url := client.GetResourceById(dashboard, id)

	req, err := http.NewRequest("DELETE", url, nil)
	response, err := client.HttpClient.Do(req)
	defer response.Body.Close()

	apiResponse := Response{response}

	err = apiResponse.HandleApiError("error deleting dashboard")

	if err != nil {
		return fmt.Errorf("error doing DELETE request: %v", err)
	}

	return nil
}

func toSlug(title string) string {
	lowerCase := strings.ToLower(title)
	return strings.Replace(lowerCase, " ", "-", -1)
}
