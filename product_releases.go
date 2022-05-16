package releases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
)

const (
	api_url = "https://api.releases.hashicorp.com/v1/releases/"
)


type ProductReleases struct {
	Releases
	Productname string
	count string
}

func (pr *ProductReleases) Fetch(ts time.Time) error {
	url := fmt.Sprintf("%s%s?after=%s&count=%s", api_url, pr.Productname, ts.Format(strfmt.ISO8601LocalTime), string(pr.count))
	fmt.Printf("url: %v\n", url)
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	rs := Releases{}
	err = json.NewDecoder(r.Body).Decode(&rs)
	if err != nil {
		return err
	}

	/*
	if len(rs) < 1 {
		return fmt.Errorf("Didn't fetched any release from: %s", url)
	}
	*/
	pr.Releases = append(pr.Releases, rs...)
	return nil
}

func (pr *ProductReleases) FetchAll() error {
	or := pr.Releases.Oldest()
	ts := or.Timestamp
	for {
		fmt.Printf("or: %v\n", or.Timestamp)
		fmt.Printf("ts: %v\n", ts)

		err := pr.Fetch(or.Timestamp)
		if err != nil {
			return err
		}
		or = pr.Releases.Oldest()
		if ts == or.Timestamp || ts.Before(or.Timestamp) {
			return nil
		}
		ts = or.Timestamp
	}
}

func New(productname string) ProductReleases{
	return ProductReleases{
		Productname: productname,
		count: "20",
	}
}

