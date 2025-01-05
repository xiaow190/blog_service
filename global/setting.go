package global

import (
	"github.com/go-programming-tour-book/blog_service/pkg/logger"
	"github.com/go-programming-tour-book/blog_service/pkg/setting"
)

//  包全局变量

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
