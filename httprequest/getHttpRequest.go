package httprequest

import(
	"strings"
	"regexp"
)

type Request struct {
	Method string
	Path string
	Protocol string
	Connection string
}

func GetHTTPRequest(buffer []byte) (bool, *Request) {
	requestString := getRequestString(buffer)
	// fmt.Println(requestString)
	// os.Stdout.Write([]byte(requestString))
	// os.Stdout.Write([]byte("\n\n"))
	requests := strings.Split(requestString, "\r\n")
	// for _, str := range requests {
	// 	os.Stdout.Write([]byte("[" + str + "]"))
	// 	// fmt.Printf("[%s]", str)
	// }
	splittedRequestString := strings.Split(requests[0], " ")
	// for _, str := range splittedRequestString {
	// 	os.Stdout.Write([]byte(str))
	// }
	var Connection string
	if 3 != len(splittedRequestString) {
		er := Request {
			Method: "no",
			Path: "no",
			Protocol: "no",
		}
		return false, &er
	}
	if !isAllowMethod(splittedRequestString[0]) {
		er := Request {
			Method: "unknown",
			Path: "no",
			Protocol: "no",
		}
		return false, &er
	}
	if !isAllowPath(splittedRequestString[1]) {
		er := Request {
			Method: "no",
			Path: "unknown",
			Protocol: "no",
		}
		return false, &er
	}
	if !isAllowProtocol(splittedRequestString[2]) {
		er := Request {
			Method: "no",
			Path: "no",
			Protocol: "unknown",
		}
		return false, &er
	}

	for _, reqElment := range requests {
		connectionReg, _ := regexp.Compile(`^Connection: keep-alive$`)
		if (connectionReg.MatchString(string(reqElment))) {
			Connection = "keep-alive"
		}
	}

	request := Request{
		Method: splittedRequestString[0],
		Path: splittedRequestString[1],
		Protocol: splittedRequestString[2],
		Connection: Connection,
	}
	return true, &request
}

func isAllowMethod(method string) (bool) {
	allowMethods := []string{"GET"}
	isAllow := false
	for _, allowMethod := range allowMethods {
		if (allowMethod == method) {
			isAllow = true
		}
	}
	return isAllow
}

func isAllowPath (path string) (bool) {
	regex := `^/.*`
	r := regexp.MustCompile(regex)
	return r.MatchString(path)
}

func isAllowProtocol(protocol string) (bool) {
	allowProtocols := []string{"HTTP/1.0", "HTTP/1.1"}
	isAllow := false
	for _, allowProtpcol := range allowProtocols {
		// fmt.Println(allowProtpcol)
		// fmt.Println(protocol)
		if (allowProtpcol == protocol) {
			isAllow = true
		}
	}
	return isAllow
}
