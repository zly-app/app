/*
-------------------------------------------------
   Author :       zlyuancn
   date：         2020/7/2
   Description :
-------------------------------------------------
*/

package core

import (
	"github.com/spf13/viper"
)

// 配置结构
type Config struct {
	// 框架配置
	Frame FrameConfig

	// 服务配置
	Services ServicesConfig

	// 插件配置
	Plugins PluginsConfig

	// 组件配置
	Components ComponentsConfig

	// 过滤器配置
	Filters FiltersConfig
}

// 配置
type IConfig interface {
	// 获取配置
	Config() *Config
	// 获取配置viper结构
	GetViper() *viper.Viper
	/*解析指定key数据到结构中
	  key 配置的key
	  outPtr 接收配置的变量
	  ignoreNotSet 如果未配置key, 则忽略, 默认为false
	*/
	Parse(key string, outPtr interface{}, ignoreNotSet ...bool) error
	/*解析组件配置, key的值为 components.{componentType}.{componentName}
	  componentType 组件类型
	  componentName 组件名
	  outPtr 接收配置的变量
	  ignoreNotSet 如果未配置key, 则忽略, 默认为false
	*/
	ParseComponentConfig(componentType ComponentType, componentName string, outPtr interface{}, ignoreNotSet ...bool) error
	/*解析插件配置, key的值为 plugins.{pluginType}
	  pluginType 插件类型
	  outPtr 接收配置的变量
	  ignoreNotSet 如果未配置key, 则忽略, 默认为false
	*/
	ParsePluginConfig(pluginType PluginType, outPtr interface{}, ignoreNotSet ...bool) error
	/*解析过滤器配置, key的值为 filters.config.{filterType}
	  filterType 过滤器类型
	  outPtr 接收配置的变量
	  ignoreNotSet 如果未配置key, 则忽略, 默认为false
	*/
	ParseFilterConfig(filterType PluginType, outPtr interface{}, ignoreNotSet ...bool) error
	/*解析服务配置, key的值为 services.{serviceType}
	  serviceType 服务类型
	  outPtr 接收配置的变量
	  ignoreNotSet 如果未配置key, 则忽略, 默认为false
	*/
	ParseServiceConfig(serviceType ServiceType, outPtr interface{}, ignoreNotSet ...bool) error
	// 检查是否存在flag, 注意: flag是忽略大小写的
	HasFlag(flag string) bool
	// 获取所有的flag, 注意: flag列表是无序的
	GetFlags() []string
	// 获取标签的值, 标签名是忽略大小写的, 标签不存在时返回空字符串
	GetLabel(name string) string
	// 获取标签数据
	GetLabels() map[string]string
	// 观察变更, 失败会fatal
	WatchKey(groupName, keyName string, opts ...ConfigWatchOption) IConfigWatchKeyObject
}

// frame配置
type FrameConfig struct {
	// debug标志
	Debug bool
	// 环境名
	Env string
	// app 名
	Name string
	// 实例名
	Instance string
	// 主动清理内存间隔时间(毫秒), <= 0 表示禁用
	FreeMemoryInterval int
	// 默认等待服务启动阶段, 等待时间(毫秒), 如果时间到未收到服务启动成功信号则将服务标记为不稳定状态然后继续开始工作(我们总不能一直等着吧)
	WaitServiceRunTime int
	// 默认服务不稳定观察时间, 等待时间(毫秒), 如果时间到仍未收到服务启动成功信号也将服务标记为启动成功
	ServiceUnstableObserveTime int
	// flag, 注意: flag是忽略大小写的
	Flags []string
	// 标签, 注意: 标签名是忽略大小写的
	Labels map[string]string
	// log配置
	Log LogConfig
	// app初始时是否打印配置
	PrintConfig bool
}

type LogConfig struct {
	Level                      string // 日志等级, debug, info, warn, error, dpanic, panic, fatal
	Json                       bool   // 启用json编码器, 输出的每一行日志转为json格式
	WriteToStream              bool   // 输出到屏幕
	WriteToFile                bool   // 日志是否输出到文件
	Name                       string // 日志文件名, 末尾会自动附加 .log 后缀
	AppendPid                  bool   // 是否在日志文件名后附加进程号
	Path                       string // 默认日志存放路径
	FileMaxSize                int    // 每个日志最大尺寸,单位M
	FileMaxBackupsNum          int    // 日志文件最多保存多少个备份, 0表示永久
	FileMaxDurableTime         int    // 文件最多保存多长时间,单位天, 0表示永久
	Compress                   bool   // 是否压缩历史日志
	TimeFormat                 string // 时间显示格式
	Color                      bool   // 是否打印彩色日志等级, 只有关闭json编码器才生效
	CapitalLevel               bool   // 是否大写日志等级
	DevelopmentMode            bool   // 开发者模式, 在开发者模式下日志记录器在写完DPanic消息后程序会感到恐慌
	ShowFileAndLinenum         bool   // 显示文件路径和行号
	ShowFileAndLinenumMinLevel string // 最小显示文件路径和行号的等级. 推荐所有等级都打印代码行, 相对于能快速定位问题来说, 这点性能损耗无关紧要
	MillisDuration             bool   // 对zap.Duration转为毫秒
}

// 服务配置
type ServicesConfig map[string]interface{}

// 插件配置
type PluginsConfig map[string]interface{}

// 组件配置
type ComponentsConfig map[string]map[string]interface{}

type FiltersConfig struct {
	Service map[string][]string
	Client  map[string]map[string][]string
	Config  map[string]interface{}
}
