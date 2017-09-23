package service_test

import (
	"reflect"
	"testing"

	"github.com/st3v/coup/service"
)

// TestRequest tests basic service request handling
func TestRequest(t *testing.T) {
	var (
		req = service.Request{
			Scooters: []int{11, 15, 13},
			C:        10,
			P:        5,
		}

		want = service.Response{NumFE: 7}
	)

	got, err := req.Handle()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected response: got %v, want %v", got, want)
	}
}
