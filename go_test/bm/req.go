package main

import (
	"context"
	"fmt"
	"go-common/library/naming/discovery"
	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/resolver"
	xtime "go-common/library/time"
	"net/http"
	"net/url"
	"time"
)

var discoveryCli *bm.Client

const EsSearchUri = "discovery://datacenter.search-plat.search-proxy/v3/search/query"

func main() {
	discoveryCli = bm.NewClient(&bm.ClientConfig{
		App: &bm.App{
			Key:    "discovery",
			Secret: "discovery",
		},
		Timeout: xtime.Duration(time.Second),
	}, bm.SetResolver(resolver.New(nil, discovery.Builder())))

	params := url.Values{}

	req, err := discoveryCli.NewRequest(http.MethodGet, EsSearchUri, "", params)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	uri := fmt.Sprintf("%s://%s%s", req.URL.Scheme, req.Host, req.URL.Path)
	err = discoveryCli.Do(context.Background(), req, nil)
	fmt.Printf("this is req uri(%v) path(%v)\n", uri, req.URL.String())
	fmt.Println("err:", err)
}
