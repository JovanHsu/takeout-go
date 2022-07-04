package nacos

//
//import (
//	"github.com/nacos-group/nacos-sdk-go/clients"
//	"github.com/nacos-group/nacos-sdk-go/common/constant"
//	"github.com/nacos-group/nacos-sdk-go/vo"
//)
//
//func main() {
//	//create clientConfig
//	clientConfig := constant.ClientConfig{
//		NamespaceId:         "e525eafa-f7d7-4029-83d9-008937f9d468", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "/tmp/nacos/log",
//		CacheDir:            "/tmp/nacos/cache",
//		LogLevel:            "debug",
//	}
//	//Another way of create clientConfig
//	clientConfig := *constant.NewClientConfig(
//		constant.WithNamespaceId("e525eafa-f7d7-4029-83d9-008937f9d468"), //When namespace is public, fill in the blank string here.
//		constant.WithTimeoutMs(5000),
//		constant.WithNotLoadCacheAtStart(true),
//		constant.WithLogDir("/tmp/nacos/log"),
//		constant.WithCacheDir("/tmp/nacos/cache"),
//		constant.WithLogLevel("debug"),
//	)
//
//	// At least one ServerConfig
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr:      "console1.nacos.io",
//			ContextPath: "/nacos",
//			Port:        80,
//			Scheme:      "http",
//		},
//		{
//			IpAddr:      "console2.nacos.io",
//			ContextPath: "/nacos",
//			Port:        80,
//			Scheme:      "http",
//		},
//	}
//	//Another way of create serverConfigs
//	serverConfigs := []constant.ServerConfig{
//		*constant.NewServerConfig(
//			"console1.nacos.io",
//			80,
//			constant.WithScheme("http"),
//			constant.WithContextPath("/nacos")
//		),
//		*constant.NewServerConfig(
//			"console2.nacos.io",
//			80,
//			constant.WithScheme("http"),
//			constant.WithContextPath("/nacos")
//		),
//	}
//
//	// Create naming client for service discovery
//	_, _ := clients.CreateNamingClient(map[string]interface{}{
//		"serverConfigs": serverConfigs,
//		"clientConfig":  clientConfig,
//	})
//
//	// Create config client for dynamic configuration
//	_, _ := clients.CreateConfigClient(map[string]interface{}{
//		"serverConfigs": serverConfigs,
//		"clientConfig":  clientConfig,
//	})
//
//	// Another way of create naming client for service discovery (recommend)
//	namingClient, err := clients.NewNamingClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfigs,
//		},
//	)
//
//	// Another way of create config client for dynamic configuration (recommend)
//	configClient, err := clients.NewConfigClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfigs,
//		},
//	)
//
//}
