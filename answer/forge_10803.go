package answer

import (
	"github.com/ggmolly/belfast/connection"
	"github.com/ggmolly/belfast/consts"
	"github.com/ggmolly/belfast/protobuf"
	"google.golang.org/protobuf/proto"
	"os"
)

// In the Azurlane code, it is called ServerInterconnection, used to switch Android/iOS servers
func Forge_SC10803_CN_JP_KR_TW(buffer *[]byte, client *connection.Client) (int, int, error) {
	belfastRegion := os.Getenv("AL_REGION")
	response := protobuf.SC_10803_CN_JP_KR_TW{
		GatewayIp:   proto.String(consts.RegionGateways[belfastRegion]),
		GatewayPort: proto.Uint32(80),
		ProxyIp:     proto.String(consts.RegionProxies[belfastRegion]),
		ProxyPort:   proto.Uint32(20000),
	}

	return client.SendMessage(10803, &response)
}
