package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/leif-sh/fog/src/utils"
)

// 自定义一个结构体，实现 gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

// 重写 Write([]byte) (int, error) 方法
func (w responseWriter) Write(b []byte) (int, error) {
	//向一个bytes.buffer中写一份数据来为获取body使用
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := utils.SugarLogger.Sync()
			if err != nil {
				utils.SugarLogger.Error("Failed append log")
			}
		}() // flushes buffer, if any

		writer := responseWriter{
			c.Writer,
			bytes.NewBuffer([]byte{}),
		}
		c.Writer = writer

		utils.SugarLogger.Infow("【request:】", "body", c.Request.Body)
		c.Next()
		utils.SugarLogger.Infow("【response:】", "body", writer.b.String())

	}
}
