package source

import (
	"fmt"
	"testing"
)

func TestGovCN_generateAjaxParams(t *testing.T) {
	govcn := NewGovCN()
	params := govcn.generateAjaxParams()
	fmt.Printf("%v\n", params)
}

func TestGovCN_Request(t *testing.T) {
	govcn := NewGovCN()
	govcn.request()
}

func TestGovCN_HighRisk(t *testing.T) {
	c := NewGovCN()
	h := c.HighRisk()
	fmt.Printf("%v\n", h)
}
