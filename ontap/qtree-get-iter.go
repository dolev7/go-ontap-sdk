package ontap

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type QtreeGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		QtreeGetIterOptions
	}
}

type QtreeGetIterOptions struct {
	MaxRecords int         `xml:"max-records,omitempty"`
	Tag        string      `xml:"tag,omitempty"`
	Query      *QtreeQuery `xml:"query,omitempty"`
}

type QtreeQuery struct {
	QtreeInfo `xml:"qtree-info,omitempty"`
}

type QtreeGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			Qtrees []QtreeInfo `xml:"qtree-info"`
		} `xml:"attributes-list"`
		NextTag string `xml:"next-tag"`
	} `xml:"results"`
}

type QtreeInfo struct {
	Qtree         string `xml:"qtree"`
	Volume        string `xml:"volume"`
	Vserver       string `xml:"vserver"`
	Status        string `xml:"status"`
	OpLocks       string `xml:"oplocks"`
	SecurityStyle string `xml:"security-style"`
	Id            string `xml:"id"`
}

func (c *Client) QtreeGetAPI(options *QtreeGetIterOptions) (*QtreeGetIterResponse, *http.Response, error) {
	if c.QtreeGetIter == nil {
		c.QtreeGetIter = &QtreeGetIter{
			Base: Base{
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.QtreeGetIter.Params.XMLName = xml.Name{Local: "qtree-list-iter"}
	}
	c.QtreeGetIter.Params.QtreeGetIterOptions = *options
	r := QtreeGetIterResponse{}
	res, err := c.QtreeGetIter.get(c.QtreeGetIter, &r)
	if err == nil && !r.Results.Passed() {
		err = fmt.Errorf("error(QtreeGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) QtreeGetIterAPI(options *QtreeGetIterOptions) (responseQtrees []*QtreeGetIterResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.QtreeGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseQtrees = append(responseQtrees, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
