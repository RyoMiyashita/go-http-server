package httprequest

import(
	"strings"
	"regexp"
	// "fmt"
)

type Request struct {
	method string
	path string
	protocol string
}

func GetHTTPRequest(buffer []byte) (bool, *Request) {
	requestString := getRequestString(buffer)
	// fmt.Println(requestString)
	splittedRequestString := strings.Split(requestString, " ")
	// for _, str := range splittedRequestString {
	// 	fmt.Printf("[%s]", str)
	// }
	if 3 != len(splittedRequestString) {
		er := Request {
			method: "no",
			path: "no",
			protocol: "no",
		}
		return false, &er
	}
	if !isAllowMethod(splittedRequestString[0]) {
		er := Request {
			method: "unknown",
			path: "no",
			protocol: "no",
		}
		return false, &er
	}
	if !isAllowPath(splittedRequestString[1]) {
		er := Request {
			method: "no",
			path: "unknown",
			protocol: "no",
		}
		return false, &er
	}
	if !isAllowProtocol(splittedRequestString[2]) {
		er := Request {
			method: "no",
			path: "no",
			protocol: "unknown",
		}
		return false, &er
	}

	request := Request{
		method: splittedRequestString[0],
		path: splittedRequestString[1],
		protocol: splittedRequestString[2],
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
