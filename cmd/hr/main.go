// hr finds and shows extended metadata about HashiCorp product releases
package main

import (
	"fmt"

	releases "github.com/smaftoul/go-hashicorp-releases"
)

func main() {
	pr := releases.New("terraforma")
	err := pr.FetchAll()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Oldest: %v : %v\n", pr.Oldest().Timestamp, pr.Oldest().Version)
	fmt.Printf("Newest: %v : %v\n", pr.Newest().Timestamp, pr.Newest().Version)
	fmt.Printf("pr.Newest().Url(): %v\n", pr.Newest().Url())
	fmt.Printf("rs.Versions(): %v\n", pr.Versions())
}
