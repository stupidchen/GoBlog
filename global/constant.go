package global

const DefaultConfigFile = "/var/log/goblog.log"
const ConfigPathEnvKey = "blogConfigPath"
const ConfigEnvKey = "blogConfig"

//root:Password@tcp(120.55.56.82:3306)/goblog?charset=utf8&&parseTime=True&loc=Local
const DefaultDBConnectionFormat = "%s:%s@tcp(%s:%s)/goblog?charset=utf8&&parseTime=True&loc=Local"