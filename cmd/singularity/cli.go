// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package main

import (
	"github.com/sylabs/singularity/cmd/singularity/cli"
	"github.com/sylabs/singularity/internal/pkg/buildcfg"
	useragent "github.com/sylabs/singularity/internal/pkg/util/user-agent"
)

func main() {
	// In cli/singularity.go
	cli.ExecuteSingularity()
}

func init() {
	useragent.InitValue(buildcfg.PACKAGE_NAME, buildcfg.PACKAGE_VERSION)
}
