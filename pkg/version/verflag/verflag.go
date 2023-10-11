// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package verflag

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/Daz-3ux/dBlog/pkg/version"
)

// used to represent the version flag value
type versionValue int

const (
	// VersionFalse not specified
	VersionFalse versionValue = iota
	// VersionTrue specified as true
	VersionTrue
	// VersionRaw specified as true and use raw format
	VersionRaw
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionFalse, "Print version info and quit")

// implement the pflag.Value interface

// IsBoolFlag returns true if the value can explain as boolean flag
func (v *versionValue) IsBoolFlag() bool {
	return true
}

// Get returns the value of the flag
func (v *versionValue) Get() interface{} {
	return v
}

// String returns the string format of the flag
func (v *versionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}

	return fmt.Sprintf("%v", *v == VersionTrue)
}

// Set sets the value of the flag
func (v *versionValue) Set(s string) error {
	if s == strRawVersion {
		*v = VersionRaw

		return nil
	}
	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}

	return err
}

// Type returns the type of the flag
func (v *versionValue) Type() string {
	return "version"
}

func VersionVar(p *versionValue, name string, value versionValue, usage string) {
	*p = value
	pflag.Var(p, name, usage)
	pflag.Lookup(name).NoOptDefVal = "true"
}

func Version(name string, value versionValue, usage string) *versionValue {
	p := new(versionValue)
	VersionVar(p, name, value, usage)

	return p
}

func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

func PrintAndExitIfRequested() {
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}
}
