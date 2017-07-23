package loging

import (
	"testing"

	logger "git.oschina.net/janpoem/go-logger"
)

func init() {
	logger.SetRollingDaily("D://logs", "mm.log")
	logger.SetLevel(logger.LOG)
}

func TestMain(t *testing.T) {
	/*
		http.HandleFunc("/", defaultIndex)
		http.HandleFunc("/login", login)
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}

		用web方式启动项目进行测试,请使用main方法,Test方法有点问题
		需要在init方法中初始化logger
		D盘下面必须先存在logs文件夹
		如果要验证按日期备份日志文件,可以启动web项目,修改系统时间,这样就能验证,生成的文件格式为mm.log.2017-07-23
	*/

	for i := 0; i < 10; i++ {
		logger.Log("name", "bush")
		logger.Log("pass", 123456)
	}

}
