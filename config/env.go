package config

import "os"

// LoadConfigMap 用来将环境变量加载到字典类型的配置中。
func LoadConfigMap(cfg map[string]string) {
	for k := range cfg {
		cfg[k] = os.Getenv(k)
	}
}
