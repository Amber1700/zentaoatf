package service

import (
	"fmt"
	serverUtils "github.com/easysoft/zentaoatf/src/server/utils/common"
	serverConst "github.com/easysoft/zentaoatf/src/server/utils/const"
	commonUtils "github.com/easysoft/zentaoatf/src/utils/common"
	configUtils "github.com/easysoft/zentaoatf/src/utils/config"
	constant "github.com/easysoft/zentaoatf/src/utils/const"
	fileUtils "github.com/easysoft/zentaoatf/src/utils/file"
	i118Utils "github.com/easysoft/zentaoatf/src/utils/i118"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"github.com/fatih/color"
	"github.com/inconshreveable/go-update"
	"github.com/mholt/archiver/v3"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var ()

type UpgradeService struct {
}

func NewUpgradeService() *UpgradeService {
	return &UpgradeService{}
}

func (s *UpgradeService) CheckUpgrade() {
	pth := vari.AgentLogDir + "version.txt"
	serverUtils.Download(serverConst.AgentUpgradeURL, pth)

	content := strings.TrimSpace(fileUtils.ReadFile(pth))
	version, _ := strconv.ParseFloat(content, 64)
	if vari.Config.Version < version {
		logUtils.PrintToWithColor(i118Utils.I118Prt.Sprintf("find_new_ver", content), color.FgCyan)

		versionStr := fmt.Sprintf("%.1f", version)
		err := s.DownloadVersion(versionStr)
		if err == nil {
			s.RestartVersion(versionStr)
		}
	}
}

func (s *UpgradeService) DownloadVersion(version string) (err error) {

	os := commonUtils.GetOs()
	if commonUtils.IsWin() {
		os = fmt.Sprintf("%s%d", os, strconv.IntSize)
	}
	url := fmt.Sprintf(serverConst.AgentDownloadURL, version, os)

	dir := vari.AgentLogDir + version
	pth := dir + ".zip"
	err = serverUtils.Download(url, pth)

	if err == nil {
		fileUtils.RmDir(dir)
		fileUtils.MkDirIfNeeded(dir)
		err = archiver.Unarchive(pth, dir)

		if err != nil {
			logUtils.PrintToWithColor(i118Utils.I118Prt.Sprintf("fail_unzip", pth), color.FgCyan)
		}
	}

	return
}

func (s *UpgradeService) RestartVersion(version string) (err error) {
	currExePath := vari.ZTFDir + constant.AppName
	bakExePath := currExePath + "_bak"
	newExePath := vari.AgentLogDir + version + constant.PthSep + constant.AppName + constant.PthSep + constant.AppName
	if commonUtils.IsWin() {
		currExePath += ".exe"
		bakExePath += ".exe"
		newExePath += ".exe"
	}
	logrus.Println(currExePath)

	rd, _ := os.Open(newExePath)
	err = update.Apply(rd, update.Options{OldSavePath: bakExePath})
	if err != nil {
		logUtils.PrintToWithColor(i118Utils.I118Prt.Sprintf("fail_upgrade",
			vari.Config.Version, version, err.Error()), color.FgRed)
	} else {
		logUtils.PrintToWithColor(i118Utils.I118Prt.Sprintf("success_upgrade",
			vari.Config.Version, version), color.FgCyan)

		// update config file
		vari.Config.Version, _ = strconv.ParseFloat(version, 64)
		configUtils.SaveConfig(vari.Config)
	}

	return
}
