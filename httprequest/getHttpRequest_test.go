package httprequest

import(
	"testing"
	"reflect"
	"fmt"
)

func TestGetHttpRequest(t *testing.T) {
	httpRequest := []byte("GET / HTTP/1.1")
	isHTTP, request := GetHTTPRequest(httpRequest)
	eisHTTP := true
	erequest := Request{
		method: "GET",
		path: "/",
		protocol: "HTTP/1.1",
	}
	if isHTTP != eisHTTP {
		t.Errorf("isHTTP actual %v\nwant %v", isHTTP, eisHTTP)
	}
	if !reflect.DeepEqual(*request, erequest) {
		t.Errorf("not equal request")
	}
}
func TestNoGetHttpRequest(t *testing.T) {
	httpRequest := []byte("GET / HTTP/1.1 awawa")
	isHTTP, request := GetHTTPRequest(httpRequest)
	eisHTTP := false
	erequest := Request{
		method: "no",
		path: "no",
		protocol: "no",
	}
	if isHTTP != eisHTTP {
		t.Errorf("isHTTP actual %v\nwant %v", isHTTP, eisHTTP)
	}
	if !reflect.DeepEqual(*request, erequest) {
		t.Errorf("not equal request")
		fmt.Printf("%+v\n", request)
		fmt.Printf("%+v\n", erequest)
	}
}

func TestIsAllowMethod(t *testing.T) {
	isAllow := isAllowMethod("GET")
	expected := true
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowMethodno(t *testing.T) {
	isAllow := isAllowMethod("AAAA")
	expected := false
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowPath(t *testing.T) {
	isAllow := isAllowPath("/index.html")
	expected := true
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowPath2(t *testing.T) {
	isAllow := isAllowPath("/")
	expected := true
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowPathno(t *testing.T) {
	isAllow := isAllowPath("a/index")
	expected := false
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowPathno2(t *testing.T) {
	isAllow := isAllowPath("")
	expected := false
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowPathno3(t *testing.T) {
	isAllow := isAllowPath("awefawef")
	expected := false
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowProtocol(t *testing.T) {
	isAllow := isAllowProtocol("HTTP/1.0")
	expected := true
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowProtocol11(t *testing.T) {
	isAllow := isAllowProtocol("HTTP/1.1")
	expected := true
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
func TestIsAllowProtocolno(t *testing.T) {
	isAllow := isAllowProtocol("AAAA")
	expected := false
	if isAllow != expected {
		t.Errorf("isAllow actual %v\nwant %v", isAllow, expected)
	}
}
