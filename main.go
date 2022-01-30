package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/sheepla/pacz/client"
)

func main() {
	var p client.Params
	p.Query = os.Args[1]
	url := client.NewURL(p)
	result := client.Execute(url)
	choices, err := find(&result)
	if err != nil {
		log.Fatal(err)
	}

	for _, idx := range choices {
		fmt.Printf("%s\n", result.Results[idx].Pkgname)
	}
}

func find(result *client.Result) ([]int, error) {
	return fuzzyfinder.FindMulti(
		result.Results,
		func(i int) string {
			if i == -1 {
				return ""
			}
			return result.Results[i].Pkgname
		},
		fuzzyfinder.WithPreviewWindow(
			func(i, width, height int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf(
					"%s/%s v%s\n\n%s\n\n%s\nlastupdate: %s\ninstalledsize: %s\n",
					result.Results[i].Repo,
					result.Results[i].Pkgname,
					result.Results[i].Pkgver,
					result.Results[i].Pkgdesc,
					result.Results[i].URL,
					humanize.Time(result.Results[i].LastUpdate),
					humanize.IBytes(uint64(result.Results[i].InstalledSize)),
				)
			},
		),
	)
}
