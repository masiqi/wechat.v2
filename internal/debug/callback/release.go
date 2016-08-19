// +build !wechatdebug

package callback

import (
	"encoding/xml"
	"io"

	"github.com/valyala/fasthttp"
)

func DebugPrintRequest(ctx *fasthttp.RequestCtx) {}

func DebugPrintPlainRequestMessage(msg []byte) {}

func XmlMarshalResponseMessage(msg interface{}) ([]byte, error) {
	return xml.Marshal(msg)
}

func XmlEncodeResponseMessage(w io.Writer, msg interface{}) error {
	return xml.NewEncoder(w).Encode(msg)
}
