package xmlStudy

import (
	"time"
	"encoding/xml"
	"fmt"
)

type XmlTime time.Time

type Teacher struct {
	Name		string		`xml:"name""`
	Age			int			`xml:"age""`
	School		string		`xml:"school"`
	Brithday	XmlTime		`xml:"brithday"`
}

func (j XmlTime)MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	//fmt.Println("这是节点名称：", start.Name.)
	if start.Name.Local == "brithday" {
		fmt.Println("这是节点名称：", start.Name.Local)

	}
	return nil
}

