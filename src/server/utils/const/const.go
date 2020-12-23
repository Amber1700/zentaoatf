package serverConst

import constant "github.com/easysoft/zentaoatf/src/utils/const"

const (
	HeartBeatInterval    = 60
	CheckUpgradeInterval = 30

	AgentRunTime = 30 * 60
	AgentLogDir  = "log-agent"

	QiNiuURL         = "https://dl.cnezsoft.com/" + constant.AppName + "/"
	AgentUpgradeURL  = QiNiuURL + "version.txt"
	AgentDownloadURL = QiNiuURL + "%s/%s/" + constant.AppName + ".zip"
)
