package main

import (
	_ "github.com/flyteorg/flyte/src/flytepropeller/plugins"

	"github.com/flyteorg/flyte/src/flytepropeller/cmd/controller/cmd"
)

func main() {
	cmd.Execute()
}
