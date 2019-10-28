package utils
import (
	"fmt"
	"testing"

)

const (
	testProposerHexAddr = "50E5518ED69AE5F0E5B737E1E9165CDD56826184"
)

func TestGenHexAddrFromPubKey(t *testing.T) {
	pubkeyVector := "hsnvalconspub1zcjduepqjlzvnup2xvanh94yf40eadzfs4e57tc63n0qlg8m6wjs8urq25esqkwnd3"
	pubkeyHex := GenHexAddrFromPubKey(pubkeyVector)
	fmt.Println(pubkeyHex)
	fmt.Println(testProposerHexAddr)
}