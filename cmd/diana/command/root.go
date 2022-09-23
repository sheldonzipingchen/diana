package command

import (
	"diana/config"
	"diana/lg"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	environment string // 运行环境
	file        string // 导入的文件名称
)

var rootCmd = &cobra.Command{
	Use:   "diana",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// init 初始化 diana 命令
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "development", "应用运行的模式设置(development, test, production)")

}

// initConfig 初始化配置
func initConfig() {
	config.Init(environment)
	lg.Init()
}

func Execute() {
	log := lg.GetLog()
	if err := rootCmd.Execute(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Command Execute Fatal.")
	}
}
