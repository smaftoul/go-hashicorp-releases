package releases

import (
	"time"
	"runtime"
)

type ReleaseStatus struct {
	Message string
	State   string
}

type Release struct {
	Builds       []Build       `json:"builds"`
	Version      string        `json:"version"`
	Timestamp    time.Time     `json:"timestamp_created"`
	IsPreRelease bool          `json:"is_prerelease"`
	Status       ReleaseStatus `json:"status"`
}

func (r Release) Url() string {
	for _, v := range r.Builds {
		if v.Arch == runtime.GOARCH && v.Os == runtime.GOOS {
			return v.Url
		}
	}
	return ""
}

