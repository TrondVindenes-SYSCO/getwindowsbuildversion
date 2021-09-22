# Go module WindowsBuildVersion

Fetches operating system version numbers using WMI

Current version is v0.1.1

# Functions
GetVersion() returns a Version struct
```
type Version struct {
	Major int
	Minor int
	Patch int
}
```