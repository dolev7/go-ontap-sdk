package ontap

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// QtreeListIter represents a request to retrieve a list of qtrees.
type QtreeListIter struct {
	Base
	Params struct {
		XMLName xml.Name
		QtreeListIterOptions
	}
}

// QtreeListIterOptions holds the parameters for the qtree-list-iter request.
type QtreeListIterOptions struct {
	MaxRecords int         `xml:"max-records,omitempty"`
	Tag        string      `xml:"tag,omitempty"`
	Query      *QtreeQuery `xml:"query,omitempty"`
}

// QtreeQuery represents the query parameters for qtree-list-iter request.
type QtreeQuery struct {
	QtreeInfo `xml:"qtree-info,omitempty"`
}

// QtreeListIterResponse represents the response from a qtree-list-iter request.
type QtreeListIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			Qtrees []QtreeInfo `xml:"qtree-info"`
		} `xml:"attributes-list"`
		NextTag string `xml:"next-tag"`
	} `xml:"results"`
}

// QtreeInfo represents information about a qtree.
type QtreeInfo struct {
	Qtree         string `xml:"qtree"`
	Volume        string `xml:"volume"`
	Vserver       string `xml:"vserver"`
	Status        string `xml:"status"`
	OpLocks       string `xml:"oplocks"`
	SecurityStyle string `xml:"security-style"`
}

// QtreeListAPI performs the API request for qtree-list-iter.
func (c *Client) QtreeListAPI(options *QtreeListIterOptions) (*QtreeListIterResponse, *http.Response, error) {
	if c.QtreeListIter == nil {
		c.QtreeListIter = &QtreeListIter{
			Base: Base{
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.QtreeListIter.Params.XMLName = xml.Name{Local: "qtree-list-iter"}
	}
	c.QtreeListIter.Params.QtreeListIterOptions = *options
	r := QtreeListIterResponse{}
	res, err := c.QtreeListIter.get(c.QtreeListIter, &r)
	if err == nil && !r.Results.Passed() {
		err = fmt.Errorf("error(QtreeListAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

// QtreeListIterAPI performs iterative API requests for qtree-list-iter until all qtrees are retrieved.
func (c *Client) QtreeListIterAPI(options *QtreeListIterOptions) (responseQtrees []*QtreeListIterResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.QtreeListAPI(options)
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
