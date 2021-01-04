package api

import (
	"fmt"
	"net/http"
)

func CreateHgClient(token string) HgClient {
	return HgClient{
		Url:        fmt.Sprintf("https://%s@api.hostedgraphite.com/api/v2", token),
		HttpClient: http.Client{},
	}
}
