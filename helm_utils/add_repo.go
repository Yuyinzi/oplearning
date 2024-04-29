package helm_utils

import (
	"fmt"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/repo"
)

func add(entry *repo.Entry) error {
	settings := cli.New()

	repoFile := settings.RepositoryConfig

	// 加载仓库配置文件
	repositories, err := repo.LoadFile(repoFile)
	// 如果文件不存在
	if err != nil {
		// 创建一个新的仓库配置对象
		repositories = repo.NewFile()
	}

	// 检查要添加的仓库是否已存在
	if repositories.Has(entry.Name) {
		fmt.Printf("仓库 %s 已存在", entry.Name)
		return nil
	}

	// 添加仓库信息到仓库配置
	repositories.Add(entry)

	// 保存更新后的仓库配置到文件
	if err = repositories.WriteFile(repoFile, 0644); err != nil {
		return fmt.Errorf("无法保存仓库配置文件", err)
	}

	fmt.Printf("成功添加仓库地址 %s", entry.Name)
	return nil
}
