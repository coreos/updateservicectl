package version

import (
	"github.com/coreos-inc/updatectl/third_party/github.com/coreos/go-semver/semver"
)

const Version = "0.1.0+git"

var SemVersion semver.Version

func init() {
	sv, err := semver.NewVersion(Version)
	if err != nil {
		panic("bad version string!")
	}
	SemVersion = *sv
}
