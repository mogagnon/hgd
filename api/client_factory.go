package api

import (
	"fmt"
	"net/http"
)

func CreateHgClientV2(token string) HgClient {
	return HgClient{
		Url:        fmt.Sprintf("https://%s@api.hostedgraphite.com/api/v2", token),
		HttpClient: http.Client{},
	}
}

func CreateHgClient(token string) HgClient {
	return HgClient{
		Url:        fmt.Sprintf("https://%s@api.hostedgraphite.com/api/v1", token),
		HttpClient: http.Client{},
	}
}
