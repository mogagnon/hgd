package api

import (
	"fmt"
	"net/http"
)

type HgClient struct {
	Url        string
	HttpClient http.Client
}

func (client *HgClient) GetResource(resource string) string {
	return fmt.Sprintf("%s%s", client.Url, fmt.Sprintf("%s/", resource))
}

func (client *HgClient) GetResourceById(resource string, id string) string {
	return fmt.Sprintf("%s%s/%s", client.Url, resource, id)
}
