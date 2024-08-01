package domain

type Opml struct {
	Head struct{} `xml:"head"`
	Body struct {
		Outline struct {
			Text  string `xml:"text,attr"`
			Items []struct {
				XmlUrl string `xml:"xmlUrl,attr"`
				Type   string `xml:"type,attr"`
				Text   string `xml:"text,attr"`
			} `xml:"outline"`
		} `xml:"outline"`
	} `xml:"body"`
}
