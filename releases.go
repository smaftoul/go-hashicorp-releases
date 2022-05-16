//TODO: implement cache
//TODO: implement count in fetch
//TODO: add missing elements in structures

package releases

import (
	"sort"
	"time"
)

type Releases []Release

// sort.Sort() interface
func (rs Releases) Len() int             { return len(rs) }
func (rs Releases) Less(r1, r2 int) bool { return rs[r1].Timestamp.Before(rs[r2].Timestamp) }
func (rs Releases) Swap(r1, r2 int)      { rs[r1], rs[r2] = rs[r2], rs[r1] }

func (rs Releases) Oldest() Release {
	if len(rs) == 0 {
		return Release{Timestamp: time.Now()}
	}
	sort.Sort(Releases(rs))
	return rs[0]
}

func (rs Releases) Newest() Release {
	if len(rs) == 0 {
		return Release{Timestamp: time.Unix(int64(0), 0)}
	}
	sort.Sort(Releases(rs))
	return rs[len(rs)-1]
}


func (rs Releases) Versions() []string{
	var versions []string
	for _, v := range rs {
		versions = append(versions, v.Version)
	}
	return versions
}
