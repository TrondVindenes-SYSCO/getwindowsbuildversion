package WindowsBuildVersion

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/StackExchange/wmi"
)

type Win32_OperatingSystem struct {
	BuildNumber string
	Version     string
}

type Version struct {
	Major int
	Minor int
	Patch int
}

func GetVersion() (*Version, error) {
	var versionInfo []Win32_OperatingSystem

	q := wmi.CreateQuery(&versionInfo, "")
	err := wmi.Query(q, &versionInfo)
	if err != nil {
		return nil, err
	}

	if len(versionInfo) != 1 {
		return nil, err
	}

	version := versionInfo[0]
	t := strings.Split(version.Version, ".")
	if len(t) != 3 {
		return nil, fmt.Errorf("version string fetched from WMI isn't in format major.minor.patch")
	}
	t1, err := strconv.Atoi(t[0])
	if err != nil {
		return nil, err
	}
	t2, err := strconv.Atoi(t[1])
	if err != nil {
		return nil, err
	}
	t3, err := strconv.Atoi(t[2])
	if err != nil {
		return nil, err
	}
	v := Version{Major: t1, Minor: t2, Patch: t3}
	return &v, nil
}
