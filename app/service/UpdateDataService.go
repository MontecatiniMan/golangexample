package service

import (
	"encoding/xml"
	"golangexample/component"
	"golangexample/entity"
	"golangexample/repository"
	"io"
	"net/http"
)

func Update() (info string, err error) {
	r, err := http.Get("https://www.treasury.gov/ofac/downloads/sdn.xml")
	if err != nil {
		return "Get data fail", err
	}
	defer r.Body.Close()
	decoder := xml.NewDecoder(r.Body)
	conn := component.DbConnection()

	for {
		token, err := decoder.Token()

		if err != nil {
			if err == io.EOF {
				break
			}

			return "Get token failed: " + err.Error(), err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "sdnEntry" {
				var sdnEntity entity.SdnEntity
				err := decoder.DecodeElement(&sdnEntity, &t)

				if err != nil {
					return "Decoding token failed: " + err.Error(), err
				}

				if sdnEntity.SdnType == "Individual" {
					repository.Insert(conn, sdnEntity)
				}
			}
		}
	}

	return "", nil
}
