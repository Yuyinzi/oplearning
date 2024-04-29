package helm_utils

type DeployRequest struct {
	RepoURL      string                 // 仓库地址
	ChartName    string                 // Chart名称
	ChartVersion string                 // Chart版本
	Namespace    string                 // 命名空间
	ReleaseName  string                 // 在kubernetes中的程序名
	Values       map[string]interface{} // values.yaml 配置文件
}
