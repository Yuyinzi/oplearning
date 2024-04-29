package helm_utils

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"
	"time"
)

func installChart(deployRequest *DeployRequest) error {

	settings := cli.New()

	actionConfig := new(action.Configuration)

	debug := func(format string, v ...interface{}) {
		if settings.Debug {
			format = fmt.Sprintf("[debug] %s\n", format)
			err := log.Output(2, fmt.Sprintf(format, v...))
			if err != nil {
				return
			}
		}
	}
	if err := actionConfig.Init(settings.RESTClientGetter(), deployRequest.Namespace, os.Getenv("HELM_DRIVER"), debug); err != nil {
		return fmt.Errorf("初始化 action 失败\n%s", err)
	}

	install := action.NewInstall(actionConfig)
	install.RepoURL = deployRequest.RepoURL
	install.Version = deployRequest.ChartVersion
	install.Timeout = time.Second * 300
	install.CreateNamespace = false
	install.Wait = false
	// kubernetes 中的配置
	install.Namespace = deployRequest.Namespace
	install.ReleaseName = deployRequest.ReleaseName

	chartRequested, err := install.ChartPathOptions.LocateChart(deployRequest.ChartName, settings)
	if err != nil {
		return fmt.Errorf("下载失败\n%s", err)
	}

	chart, err := loader.Load(chartRequested)
	if err != nil {
		return fmt.Errorf("加载失败\n%s", err)
	}

	_, err = install.Run(chart, nil)
	if err != nil {
		return fmt.Errorf("执行失败\n%s", err)
	}

	return nil
}
