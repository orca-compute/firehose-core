package main

import (
	firecore "github.com/streamingfast/firehose-core"
	pbbstream "github.com/streamingfast/pbgo/sf/bstream/v1"
)

func main() {
	firecore.Main(&firecore.Chain[*pbbstream.Block]{
		ShortName:            "core",      //used to compose the binary name
		LongName:             "CORE",      //only used to compose cmd title and description
		ExecutableName:       "fire-core", //only used to set default value of reader-node-path, we should not provide a default value anymore ...
		FullyQualifiedModule: "github.com/streamingfast/firehose-core",
		Version:              version,

		Protocol:        "NEA",
		ProtocolVersion: 1,

		ConsoleReaderFactory: supervisor.NewConsoleReader,

		Tools: &firecore.ToolsConfig[*pbbstream.Block]{},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
