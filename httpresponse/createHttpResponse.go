package httpresponse

import(
	"bytes"
	"strconv"
)

type Header struct {
	ContentType string `json:"Content-Type"`
}

func CreateHttpResponse(protocol string, status int, headers []string, content []byte) ([]byte, error) {
	var resHeaderString string
	var resHeaderBytes []byte
	resHeaderString += protocol + " " + strconv.Itoa(status) + "\n"
	for _, element := range headers {
		resHeaderString += element
		resHeaderString += "\n"
	}
	resHeaderBytes = []byte(resHeaderString)
	resArray := [][]byte{resHeaderBytes, content}
	// resArray[0] = resHeaderBytes
	// resArray[1] = content
	res := bytes.Join(resArray, []byte("\n"))
	return res, nil
}
