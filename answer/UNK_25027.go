package answer

import (
	"fmt"

	"github.com/bettercallmolly/belfast/connection"
	"github.com/bettercallmolly/belfast/logger"

	"github.com/bettercallmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
)

// A get with a type?
func UNK_25027(buffer *[]byte, client *connection.Client) (int, int, error) {
	var packet protobuf.CS_25026
	err := proto.Unmarshal(*buffer, &packet)
	if err != nil {
		return 0, 25027, err
	}
	logger.LogEvent("Client", "CS_25026", fmt.Sprintf("client asked for type=%d", packet.GetType()), logger.LOG_LEVEL_DEBUG)

	response := protobuf.SC_25027{
		Level: proto.Uint32(1),
		Exp:   proto.Uint32(0),
		Clean: proto.Uint32(0),
	}
	// Answer with default valid SC_25027
	return client.SendMessage(25027, &response)
}
