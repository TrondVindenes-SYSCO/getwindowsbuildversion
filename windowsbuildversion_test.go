package WindowsBuildVersion

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	version, err := GetVersion()
	if version == nil || err != nil {
		t.Fatalf("Failed collecting Version struct, %q, %s", version, err)
	}
}
