package elasticsearch

import (
	"fmt"
	"github.com/system-server2025/global"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)


func InitEs() *elasticsearch.Client {
	esURL := fmt.Sprintf("http://%s:%d", global.GVA.Config.ElasticsearchConfig.Host, global.GVA.Config.ElasticsearchConfig.Port)
	var err error
	// 根据配置初始化Elasticsearch客户端
	ESClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{esURL},
	})
	if err != nil {
		log.Fatalf("无法初始化Elasticsearch客户端: %v", err)
	}
	// 如果启用认证，进行认证相关的设置（这里只是简单示例，实际可能更复杂）
	if global.GVA.Config.ElasticsearchConfig.EnableAuth {
		// 可以在这里添加认证逻辑，如设置BasicAuth等
		// 例如：
		// ESClient.Transport.(*http.Transport).TLSClientConfig = &tls.Config{
		//     InsecureSkipVerify: true,
		// }
		// ESClient.SetBasicAuth(cfg.Username, cfg.Password)
		log.Fatal("认证逻辑尚未完全实现")
	}
	// 测试连接
	res, err := ESClient.Info()
	if err != nil {
		log.Fatalf("无法连接Elasticsearch: %v", err)
	}
	defer res.Body.Close()
	// 检查连接状态
	if res.IsError() {
		log.Fatalf("Elasticsearch返回错误: %v", res.String())
	}
	return ESClient
}
