package httprequest

import (
	"testing"
  )

  func TestGetRequestString(t *testing.T) {
	buffer := []byte("hello")
	actual := getRequestString(buffer)
	expected := "hello"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
  }
