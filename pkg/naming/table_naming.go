package naming

import sca_base_module_config "github.com/scagogogo/sca-base-module-config"

// TableName 为所有的表统一命名
func TableName(tableName string) string {
	if sca_base_module_config.Config != nil {
		// 如果需要在同一个库测试的话，则给表名称拼接一个前缀，但这不是推荐的做法
		return sca_base_module_config.Config.GetString("database.table-name-prefix") + tableName
	} else {
		return tableName
	}
}
