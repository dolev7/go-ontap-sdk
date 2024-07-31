package ontap

import (
	"encoding/xml"
)

type QtreeGetIter struct {
	Base
	Params struct {
		XMLName xml.Name `xml:"qtree-get-iter"`
	} `xml:"parameters"`
}

type QtreeGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		SingleResultBase
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
}
