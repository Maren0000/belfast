package answer

import (
	"github.com/ggmolly/belfast/connection"

	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func GuildGetAssaultFleetCommandResponse(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_61012{
		Result: proto.Uint32(0),
	}
	return client.SendMessage(61012, &response)
}
