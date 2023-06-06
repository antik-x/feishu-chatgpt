package initialization

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
)

var larkClient *lark.Client

func LoadLarkClient(config Config) {
	larkClient = lark.NewClient(config.FeishuAppId, config.FeishuAppSecret, lark.WithLogLevel(larkcore.LogLevelDebug))
}

func GetLarkClient() *lark.Client {
	return larkClient
}
