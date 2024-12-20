package heavycheck

import (
	"github.com/Fantom-foundation/lachesis-base/inter/dag"

	"github.com/WlinkNET/xpense_chain/inter"
)

type EventsOnly struct {
	*Checker
}

func (c *EventsOnly) Enqueue(e dag.Event, onValidated func(error)) error {
	return c.Checker.EnqueueEvent(e.(inter.EventPayloadI), onValidated)
}
