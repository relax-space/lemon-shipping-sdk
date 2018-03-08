package best

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
)

func toBase64(dto interface{}) (data string, err error) {
	da, err := json.Marshal(dto)
	if err != nil {
		return
	}
	encodeData := url.QueryEscape(string(da))
	data = base64.StdEncoding.EncodeToString([]byte(encodeData))
	return
}

const (
	SUC = "SUC" //success
	E01 = "E01" //system error,can re-try
	E02 = "E02" //bad request format
	E03 = "E03" //message from shipping
	E04 = "E04" //bad response format
)
