// 初始化包

package init

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"example.com/ningneng/internal/config"
	"example.com/ningneng/internal/zap"
	"example.com/ningneng/pkg/global"
)

func init() {
	var err error

	// 获取项目根目录路径
	if global.RootPath, err = getRootPath(); err != nil {
		log.Fatalf("获取Root Path错误：%v", err)
	}

	println("RootPath:", global.RootPath)

	// 加载配置文件
	config.Init()

	// 时区
	if global.Location, err = time.LoadLocation(
		global.Viper.GetString("application.time_zone")); err != nil {
		log.Fatalf("时区错误：%v", err)
	}

	// 日志
	if global.Logger, err = zap.NewLogger(false); err != nil {
		log.Fatalf("日志初始化错误：%v", err)
	}
}

func getRootPath() (root string, err error) {
	root, err = getCurrentAbPathByExecutable()
	if strings.Contains(root, getTemDir()) {
		return getCurrentAbPathByCaller(), nil
	}
	return root, err
}

func getTemDir() string {
	var dir string

	switch runtime.GOOS {
	case "linux", "darwin":
		dir = os.Getenv("TMPDIR")
	case "windows":
		dir = os.Getenv("TEMP")
		if dir == "" {
			dir = os.Getenv("TMP")
		}
	}

	if dir == "" {
		dir = "/tmp"
	}

	res, _ := filepath.EvalSymlinks(dir)

	return res
}

func getCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(filepath.Dir(exePath))
}

func getCurrentAbPathByCaller() (root string) {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		root = path.Dir(path.Dir(filename))
	}
	return
}
