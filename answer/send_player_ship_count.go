package answer

import (
	"github.com/ggmolly/belfast/consts"
	"os"
	"time"

	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func SendPlayerShipCount(buffer *[]byte, client *connection.Client) (int, int, error) {
	belfastRegion := os.Getenv("AL_REGION")
	answer := protobuf.SC_11002{
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(consts.Monday_0OclockTimestamps[belfastRegion]),
		ShipCount:               proto.Uint32(uint32(len(client.Commander.Ships))),
	}
	client.Server.JoinRoom(client.Commander.RoomID, client)
	return client.SendMessage(11002, &answer)
}
