package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Params struct {
	Query      string
	Name       string
	Desc       string
	Repo       string
	Arch       string
	Maintainer string
	Packager   string
	Flagged    bool
}

type Result struct {
	Limit    int64 `json:"limit"`
	NumPages int64 `json:"num_pages"`
	Page     int64 `json:"page"`
	Results  []struct {
		Arch           string    `json:"arch"`
		BuildDate      string    `json:"build_date"`
		Checkdepends   []string  `json:"checkdepends"`
		CompressedSize int       `json:"compressed_size"`
		Conflicts      []string  `json:"conflicts"`
		Depends        []string  `json:"depends"`
		Epoch          int       `json:"epoch"`
		Filename       string    `json:"filename"`
		FlagDate       time.Time `json:"flag_date"`
		Groups         []string  `json:"groups"`
		InstalledSize  int       `json:"installed_size"`
		LastUpdate     time.Time `json:"last_update"`
		Licenses       []string  `json:"licenses"`
		Maintainers    []string  `json:"maintainers"`
		Makedepends    []string  `json:"makedepends"`
		Optdepends     []string  `json:"optdepends"`
		Packager       string    `json:"packager"`
		Pkgbase        string    `json:"pkgbase"`
		Pkgdesc        string    `json:"pkgdesc"`
		Pkgname        string    `json:"pkgname"`
		Pkgrel         string    `json:"pkgrel"`
		Pkgver         string    `json:"pkgver"`
		Provides       []string  `json:"provides"`
		Replaces       []string  `json:"replaces"`
		Repo           string    `json:"repo"`
		URL            string    `json:"url"`
	} `json:"results"`
	Valid   bool  `json:"valid"`
	Version int64 `json:"version"`
}

func NewURL(params Params) string {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "www.archlinux.org"
	u.Path = "packages/search/json"
	q := u.Query()
	if params.Query != "" {
		q.Set("q", params.Query)
	}
	if params.Name != "" {
		q.Set("name", params.Name)
	}
	if params.Desc != "" {
		q.Set("desc", params.Desc)
	}
	if params.Arch != "" {
		q.Set("arch", params.Arch)
	}
	if params.Maintainer != "" {
		q.Set("arch", params.Maintainer)
	}
	if params.Packager != "" {
		q.Set("arch", params.Packager)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func Execute(url string) Result {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	return result
}
