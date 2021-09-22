# Go module WindowsBuildVersion

Fetches operating system version numbers using WMI

# Functions
GetVersion() returns a Version struct
```
type Version struct {
	Major int
	Minor int
	Patch int
}
```