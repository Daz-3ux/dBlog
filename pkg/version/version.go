// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	// GitVersion is a semantic version number
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate format is ISO8601
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit is the commit hash
	GitCommit = "$Format:%H$"
	// GitTreeState is the state of Git repository during the build
	// 'clean' means  no uncommitted changes
	// 'dirty' means there are uncommitted changes
	GitTreeState = ""
)

// Info include the version info
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns the readable version info
func (info Info) String() string {
	if s, err := info.Text(); err == nil {
		return string(s)
	}

	return info.GitVersion
}

// ToJson returns the json format version info
func (info Info) ToJson() string {
	s, _ := json.Marshal(info)

	return string(s)
}

// Text encodes version info in UTF-8 formatted text and return it
func (info Info) Text() ([]byte, error) {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)

	return table.Bytes(), nil
}

// Get return comprehensive version info
func Get() Info {
	return Info{
		// set by ldflags
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
