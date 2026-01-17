package runtimemeta

import (
	"fmt"
	"runtime/debug"
	"time"
)

type RuntimeMeta struct {
	HostName  string
	Path      string
	StartTime time.Time

	Version       string
	Revision      string
	RevisionShort string
	LastCommit    time.Time
	DirtyBuild    bool

	GoVersion string
	GoArch    string
	GoOS      string
}

func GetRuntimeMeta(version string, hostName string) (*RuntimeMeta, error) {
	rm := RuntimeMeta{
		HostName:  hostName,
		StartTime: time.Now(),
		Version:   version,
	}

	if rm.Version == "" {
		rm.Version = "latest"
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return &rm, fmt.Errorf("unable to get debug build info")
	}

	rm.GoVersion = info.GoVersion
	rm.Path = info.Path

	for _, kv := range info.Settings {
		if kv.Value == "" {
			continue
		}
		switch kv.Key {
		case "GOARCH":
			rm.GoArch = kv.Value
		case "GOOS":
			rm.GoOS = kv.Value
		case "vcs.revision":
			rm.Revision = kv.Value
		case "vcs.time":
			rm.LastCommit, _ = time.Parse(time.RFC3339, kv.Value)
		case "vcs.modified":
			rm.DirtyBuild = kv.Value == "true"
		}
	}

	if len(rm.Revision) >= 7 {
		rm.RevisionShort = rm.Revision[0:7]
	}

	return &rm, nil
}
