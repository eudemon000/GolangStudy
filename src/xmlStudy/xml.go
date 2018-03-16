package xmlStudy

import (
	"encoding/xml"
	"fmt"
)

func XMLMarshal(t *Teacher) {
	b, err := xml.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("转化后的XML：", string(b))

}
