package integration

import (
	"path"

	"github.com/WlinkNET/xpense_chain/utils"
)

func isInterrupted(chaindataDir string) bool {
	return utils.FileExists(path.Join(chaindataDir, "unfinished"))
}
