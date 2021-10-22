package rss2

import (
	"encoding/xml"
	"fmt"
)

// Cloud represents a Channel's cloud element. All attributes must be
// present.
type Cloud struct {
	XMLName           xml.Name `xml:"cloud"`
	Domain            string   `xml:"domain,attr"`
	Port              int      `xml:"port,attr"`
	Path              string   `xml:"path,attr"`
	RegisterProcedure string   `xml:"registerProcedure,attr"`
	Protocol          string   `xml:"protocol,attr"`
}

// NewCloud creates a new Cloud element.
func NewCloud(domain string, port int, path, rp, protocol string) (*Cloud, error) {
	if len(domain) == 0 || len(path) == 0 || len(rp) == 0 || len(protocol) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewCloud()`)
	}
	return &Cloud{
		XMLName:           xml.Name{Local: `cloud`},
		Domain:            domain,
		Port:              port,
		Path:              path,
		RegisterProcedure: rp,
		Protocol:          protocol,
	}, nil
}
