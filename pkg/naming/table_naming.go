package naming

import sca_base_module_config "github.com/scagogogo/sca-base-module-config"

// TableName 为所有的表统一命名
func TableName(tableName string) string {
	if sca_base_module_config.Config != nil {
		return sca_base_module_config.Config.GetString("database.table-name-prefix") + tableName
	} else {
		return tableName
	}
}
