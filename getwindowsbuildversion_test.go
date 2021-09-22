package GetWindowsBuildVersion

import (
	"testing"
)

func TestGetWindowBuildVersion(t *testing.T) {
	version, err := GetWindowBuildVersion()
	if version == nil || err != nil {
		t.Fatalf("Failed collecting Version struct, %q, %s", version, err)
	}
}
