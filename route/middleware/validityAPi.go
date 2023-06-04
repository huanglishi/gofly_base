package middleware

import (
	"gofly/global"
	"gofly/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 验证接口合法性
func ValidityAPi() gin.HandlerFunc {
	return func(c *gin.Context) {
		//加载配置
		conf := global.App.Config
		var apisecret = conf.App.Apisecret
		encrypt := c.Request.Header.Get("verify-encrypt")
		verifytime := c.Request.Header.Get("verify-time")
		mdsecret := utils.Md5(apisecret + verifytime)
		// 验证-根目录
		var NoVerifyRoot_arr []string
		if global.App.Config.App.NoVerifyRoot != "" {
			NoVerifyRoot_arr = strings.Split(global.App.Config.App.NoVerifyRoot, `,`)
		} else {
			NoVerifyRoot_arr = make([]string, 0)
		}
		// 验证-具体路径
		var noValidity []string
		if global.App.Config.App.NoValidity != "" {
			noValidity = strings.Split(global.App.Config.App.NoValidity, `,`)
		} else {
			noValidity = make([]string, 0)
		}
		rootPath := strings.Split(c.Request.URL.Path, "/")
		if encrypt == "" || mdsecret != encrypt {
			if (len(rootPath) > 2 && IsContain(NoVerifyRoot_arr, rootPath[1])) || IsContain(noValidity, c.Request.URL.Path) || strings.Contains(c.Request.URL.Path, "/common/uploadfile/get_image") { //过滤附件访问接口
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    1,
					"message": "您的请求不合法，请按规范请求数据!",
					"result":  nil,
				})
				c.Abort()
			}
		} else {
			c.Next()
		}
	}
}
