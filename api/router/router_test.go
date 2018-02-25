package router

import (
	"testing"

	storeTesting "github.com/8tomat8/Qm9yeXMtSHVsaWk-/store/testing"
)

func TestNewRouter(t *testing.T) {
	// This test could be extended to check if all routes was registered,
	// but we would need to store uor routes in some common structure
	// So far it checks only if there are at least one registered route in router
	r := NewRouter(&storeTesting.StorageMock{})

	routes := r.Routes()

	if len(routes) == 0 {
		t.Error("NewRouted was not ")
	}
}
