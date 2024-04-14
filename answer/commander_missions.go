package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/protobuf"
)

func CommanderMissions(buffer *[]byte, client *connection.Client) (int, int, error) {
	var response protobuf.SC_20001
	return client.SendMessage(20001, &response)
}
