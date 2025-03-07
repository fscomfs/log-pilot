package quota

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	docker "github.com/docker/docker/client"
	"github.com/fscomfs/cvmart-log-pilot/config"
	"os"
	"testing"
)

func TestSetDirQuota(t *testing.T) {
	config.ParseFromFile("/data/daemon-config")
	fileInfo, err := os.Stat(config.GlobConfig.HostTempDataPath)
	if os.IsNotExist(err) {
		os.MkdirAll(config.GlobConfig.HostTempDataPath, 0777)
	} else if !fileInfo.IsDir() {
		log.Errorf("path  %s is not dir", err.Error())
		return
	}
	ctl, err := NewControl("", config.GlobConfig.HostTempDataPath)
	if err != nil {
		log.Errorf("NewControl err=%s", err.Error())
		return
	}

	e := ctl.SetDirQuota("/home/data/tempdisk/user-1", Quota{1024 * 1024 * 1024})
	if e != nil {
		log.Errorf("Set Dir quotal err %s", err.Error())
	}

}

func TestSGetDirQuota(t *testing.T) {
	config.ParseFromFile("/etc/cvmart/daemon-config.json")
	fileInfo, err := os.Stat(config.GlobConfig.HostTempDataPath)
	if os.IsNotExist(err) {
		os.MkdirAll(config.GlobConfig.HostTempDataPath, 0777)
	} else if !fileInfo.IsDir() {
		log.Errorf("path  %s is not dir", err.Error())
		return
	}
	ctl, err := NewControl("", config.GlobConfig.HostTempDataPath)
	if err != nil {
		log.Errorf("NewControl err=%s", err.Error())
		return
	}

	quotaInfo, e := ctl.GetDirQuota("/home/data/tempdisk/user-1")
	if e != nil {
		log.Errorf("Get Dir quota err %+v", e.Error())
	}
	res, e := json.Marshal(quotaInfo)
	if e == nil {
		log.Infof("Get quota info %+v", string(res))
	}

	quotaInfo2, e := ctl.GetNodeSpaceInfo("/testQuotaR/testQuota/tmpdisk/user-1")
	res2, e := json.Marshal(quotaInfo2)
	if e == nil {
		log.Infof("Get Node Space quota info %+v", string(res2))
	}

}

func TestReleaseDir(t *testing.T) {
	config.ParseFromFile("/data/daemon-config")
	fileInfo, err := os.Stat(config.GlobConfig.HostTempDataPath)
	if os.IsNotExist(err) {
		os.MkdirAll(config.GlobConfig.HostTempDataPath, 0777)
	} else if !fileInfo.IsDir() {
		log.Errorf("path  %s is not dir", err.Error())
		return
	}
	ctl, err := NewControl("", config.GlobConfig.HostTempDataPath)
	if err != nil {
		log.Errorf("NewControl err=%s", err.Error())
		return
	}

	e := ctl.ReleaseDir("/home/data/tempdisk/user-1")
	if e != nil {
		log.Errorf("Set Dir quotal err %s", err.Error())
	}

}

func TestGetRealPath(t *testing.T) {
	linkPath := "/tmp/A/B/222"

	rel, err := FindRealPath("", linkPath)
	if err != nil {
		log.Errorf(rel, err)
	} else {
		log.Printf(rel)
	}
}

func TestGetMountInfo(t *testing.T) {
	s, e := GetXFSMount("")
	if e == nil {
		log.Printf("mount info %+v", s)
	}
}

func TestImageQuotaInfo(t *testing.T) {
	config.ParseFromFile("/etc/cvmart/daemon-config.json")
	fileInfo, err := os.Stat(config.GlobConfig.HostTempDataPath)
	if os.IsNotExist(err) {
		os.MkdirAll(config.GlobConfig.HostTempDataPath, 0777)
	} else if !fileInfo.IsDir() {
		log.Errorf("path  %s is not dir", err.Error())
		return
	}
	ctl, err := NewControl("", config.GlobConfig.HostTempDataPath)
	if err != nil {
		log.Errorf("NewControl err=%s", err.Error())
		return
	}
	dockerClient, _ := docker.NewEnvClient()
	quotaInfo2, e := ctl.GetImageDiskQuotaInfo("192.168.1.76:8099/evtrain/cvmart-daemon-arm64:v1.0", dockerClient)
	res2, e := json.Marshal(quotaInfo2)
	if e == nil {
		log.Infof("GetImageDiskQuotaInfo quota info %+v", string(res2))
	}

}
