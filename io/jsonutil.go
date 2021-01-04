package io

import (
	"encoding/json"
	"fmt"
	"hgd/api"
	"io/ioutil"
)

func SaveDashboard(dashboard *api.DashboardTemplate, output string) error {
	var definition = dashboard.Dashboard

	data, _ := json.Marshal(definition)
	err := ioutil.WriteFile(output, data, 0644)

	if err != nil {
		return fmt.Errorf("unable to write to the output file %s: %v", output, err)
	}

	return nil
}

func LoadDashboard(path string) (*api.DashboardTemplate, error) {
	var result api.DashboardDefinition
	file, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("unable to load file %s: %v", path, err)
	}

	err = json.Unmarshal(file, &result)

	if err != nil {
		return nil, fmt.Errorf("error to read json file %s: %v", path, err)
	}

	return &api.DashboardTemplate{Dashboard: result}, nil
}
