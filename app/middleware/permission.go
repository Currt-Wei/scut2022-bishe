package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scut2022-bishe/util/casbin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
		claims := c.MustGet("claims").(*CustomClaims)

		if ok, err := casbin.CasbinObj.Enforcer.Enforce(claims.Email, c.Request.URL, c.Request.Method); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusOK,
				"data":   err,
				"msg":    "ok",
			})
			c.Abort()
			return
		} else if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"data":   "没有权限",
				"msg":    "",
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
