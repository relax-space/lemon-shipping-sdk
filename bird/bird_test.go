package bird

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/relax-space/go-kit/test"
)

func Test_signBird(t *testing.T) {
	d, err := signBird("{'OrderCode':'','ShipperCode':'SF','LogisticCode':'118954907573'}56da2cf8-c8a2-44b2-b6fa-476cd7d1ba17")
	test.Ok(t, err)
	nd, err := url.QueryUnescape(d)
	fmt.Println(nd)
	test.Ok(t, err)
	test.Equals(t, "OWFhM2I5N2ViM2U2MGRkMjc4YzU2NmVlZWI3ZDk0MmE=", nd)
}

func Test_Query(t *testing.T) {
	custDto := &ReqCustomerDto{
		Url:    os.Getenv("BIRD_URL"),
		ApiKey: os.Getenv("BIRD_APIKEY"),
	}
	reqDto := &ReqQueryDto{
		ReqBase: &ReqBase{
			EBusinessId: os.Getenv("BIRD_EBusinessId"),
			RequestType: "1002",
			DataType:    "2",
		},
		RequestData: &ReqQueryDataDto{
			OrderCode:    "",
			ShipperCode:  "SF",
			LogisticCode: "118461988807",
			IsHandleInfo: "0",
		},
	}
	_, _, respDto, err := Query(reqDto, custDto)
	spew.Dump(respDto)
	test.Ok(t, err)
}
