package antport

import (
	"github.com/pjvds/antport/ant"
	"github.com/pjvds/antport/hardware"
	"log"
	"testing"
)

func TestOpenContext(t *testing.T) {
	device, antContext := GetSingleChannelOrFail(t)
	defer antContext.Close()
	defer device.Close()

	ctx := ant.CreateAntContext(device)
	ctx.HardResetSystem(5)
	ctx.Init()
	defer ctx.Close()

	fs := fs.NewAntFsContext(ctx)
	fs.OpenAntsFsSearchChannel()

	reply, err := ctx.ReceiveReply()
	log.Printf("Reply: %s", reply)
	log.Printf("error: %s", err)
}
