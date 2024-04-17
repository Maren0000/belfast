package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

func UpdateCommonFlagCommand(buffer *[]byte, client *connection.Client) (int, int, error) {
	response := protobuf.SC_11020{
		Result: proto.Uint32(0),
	}
	return client.SendMessage(11020, &response)
}
