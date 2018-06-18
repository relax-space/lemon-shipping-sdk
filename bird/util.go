package bird

import (
	"encoding/base64"
	"net/url"

	"github.com/relax-space/go-kit/sign"
)

func signBird(text string) (string, error) {
	md5Data, err := sign.GetMD5Hash(text, true)
	if err != nil {
		return "", err
	}
	da := base64.URLEncoding.EncodeToString([]byte(md5Data))
	encodeData := url.QueryEscape(da)
	return encodeData, nil
}

const (
	SUC = "SUC" //success
	E01 = "E01" //system error,can re-try
	E02 = "E02" //bad request format
	E03 = "E03" //message from shipping
	E04 = "E04" //bad response format
)
