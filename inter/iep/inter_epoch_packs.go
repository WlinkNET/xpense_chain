package iep

import (
	"github.com/WlinkNET/xpense_chain/inter"
	"github.com/WlinkNET/xpense_chain/inter/ier"
)

type LlrEpochPack struct {
	Votes  []inter.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
