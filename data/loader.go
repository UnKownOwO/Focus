package data

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

type LoadPriority int

const (
	LOWEST LoadPriority = iota
	LOW
	NORMAL
	HIGH
	HIGHEST
)

type ResLoader struct {
	Path     []string
	Handle   func([][]byte, []string)
	Regexp   string
	Priority LoadPriority
}

var ResLoaders = make([]*ResLoader, 0)

func LoadAll() {
	log.Println("正在载入资源文件...")
	start := time.Now()

	LoadResources()
	LoadGameDepot()

	log.Println(fmt.Sprintf("资源文件加载完毕 (%s)", time.Now().Sub(start)))
}

func LoadResources() {
	priLoaders := make(map[LoadPriority][]*ResLoader)
	// 按照优先级分开
	for _, loader := range ResLoaders {
		_, ok := priLoaders[loader.Priority]
		if !ok {
			priLoaders[loader.Priority] = make([]*ResLoader, 0, len(ResLoaders))
		}
		priLoaders[loader.Priority] = append(priLoaders[loader.Priority], loader)
	}

	var wg sync.WaitGroup
	// 从优先级最高的开始遍历
	for i := HIGHEST; i > LOWEST-1; i-- {
		// 获取此分级的所有加载器
		loaders, ok := priLoaders[i]
		if !ok {
			continue
		}
		// 执行此分级的每个加载器
		for _, loader := range loaders {
			loader := loader

			wg.Add(1)
			go func() {
				defer wg.Done()
				loader.Handle(ReadPathFile(loader.Path, loader.Regexp))
			}()
		}
		wg.Wait()
	}
}

func ReadPathFile(paths []string, expr string) ([][]byte, []string) {
	files := make([][]byte, 0, len(paths))
	match := make([]string, 0, len(paths))

	for _, path := range paths {
		var filePaths []string
		// 判断是否为文件夹
		// Bin资源为文件夹内全部
		// Excel资源为单文件
		if IsDir(path) {
			// 读取路径下的所有文件
			dir, err := os.ReadDir(path)
			if err != nil {
				log.Println("[Loader] 读取文件夹错误, path:", path, "err:", err)
				return nil, nil
			}
			filePaths = make([]string, 0, len(dir))
			// 遍历每个文件
			for _, fileInfo := range dir {
				// 跳过文件夹
				if fileInfo.IsDir() {
					continue
				}
				filePaths = append(filePaths, path+fileInfo.Name())
			}
		} else {
			filePaths = []string{path}
		}
		// 读取每一个文件的内容
		for _, filePath := range filePaths {
			// 匹配文件路径正则 不符合则跳过
			if expr != "" {
				r, _ := regexp.Compile(expr)
				if matchArr := r.FindStringSubmatch(filePath); len(matchArr) > 0 {
					match = append(match, matchArr[len(matchArr)-1])
				} else {
					continue
				}
			}
			file, err := os.ReadFile(filePath)
			if err != nil {
				log.Println("[Loader] 读取文件错误, filePath:", filePath, "err:", err)
				return nil, nil
			}
			files = append(files, file)
		}
	}
	return files, match
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
