package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// RepoInfo 结构体用于存储仓库信息，与JSON字段对应
type RepoInfo struct {
	URL    string `json:"url"`    // Git仓库URL（需包含.git后缀）
	Branch string `json:"branch"` // 要克隆的分支名称
	Path   string `json:"path"`   // 仓库的子路径（可选）
}

// Config 结构体用于存储整个配置，包括统一的项目目录
type Config struct {
	BasePath string     `json:"basePath"` // 统一的项目目录路径
	Repos    []RepoInfo `json:"repos"`    // 仓库列表
}

func main() {
	// 提示用户输入 JSON 文件路径
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入 JSON 文件路径: ")
	repoInfoFile, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("读取输入错误: ", err)
	}
	// 去除换行符
	repoInfoFile = strings.TrimSpace(repoInfoFile)

	// 读取并解析JSON配置文件
	repoInfoData, err := os.ReadFile(repoInfoFile)
	if err != nil {
		log.Fatalf("读取文件错误: %v", err)
	}

	// 反序列化JSON数据到配置结构体
	var config Config
	if err := json.Unmarshal(repoInfoData, &config); err != nil {
		log.Fatalf("解析 JSON 错误: %v\n请检查JSON格式是否正确（如逗号、引号等）", err)
	}

	// 确保统一的项目目录存在，如果不存在则创建
	if config.BasePath == "" {
		log.Fatal("JSON 配置中未指定 basePath")
	}

	err = os.MkdirAll(config.BasePath, os.ModePerm)
	if err != nil {
		log.Fatalf("创建统一项目目录 %s 时出错: %v", config.BasePath, err)
	}

	// 遍历所有仓库配置进行克隆操作
	for _, repo := range config.Repos {
		var fullPath string

		// 如果 path 字段存在且不为空，则使用 path 作为克隆路径
		if repo.Path != "" {
			fullPath = filepath.Join(config.BasePath, repo.Path)
		} else {
			// 如果 path 字段不存在或为空，则使用仓库名称作为子目录
			repoName := strings.TrimSuffix(
				filepath.Base(repo.URL), // 从URL获取最后部分
				".git",                  // 去除.git后缀
			)
			fullPath = filepath.Join(config.BasePath, repoName)
		}

		// 确保目标目录的父目录存在
		parentDir := filepath.Dir(fullPath)
		if parentDir != "" {
			err := os.MkdirAll(parentDir, os.ModePerm)
			if err != nil {
				log.Printf("创建父目录 %s 时出错: %v", parentDir, err)
				continue // 跳过当前仓库，继续下一个
			}
		}

		// 打印当前操作信息
		fmt.Printf("正在克隆 %s 的 %s 分支到 %s...\n", repo.URL, repo.Branch, fullPath)

		// 创建git命令：克隆指定分支到指定目录
		cmd := exec.Command("git", "clone",
			"-b", repo.Branch, // -b 表示指定分支
			repo.URL, // 仓库URL
			fullPath, // 目标目录路径
		)

		// 实时显示命令输出（将git的输出直接连接到标准IO）
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// 执行命令并处理错误
		if err := cmd.Run(); err != nil {
			// 克隆失败时记录错误（但继续执行后续仓库）
			log.Printf("克隆 %s 时出错: %v（可能原因：分支不存在、目录已存在或没有权限）", repo.URL, err)
		} else {
			// 成功时显示完成信息
			fmt.Printf("成功克隆 %s 的 %s 分支到 %s\n", repo.URL, repo.Branch, fullPath)
		}
	}
}
