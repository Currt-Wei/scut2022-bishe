package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scut2022-bishe/util/casbin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
		claims := c.MustGet("claims").(*CustomClaims)

		list := casbin.CasbinObj.Enforcer.GetAllNamedObjects("p")
		for _, vlist := range list {
			for _, v := range vlist {
				fmt.Printf("value: %s, ", v)
			}
			fmt.Println()
		}
		ok, _ := casbin.CasbinObj.Enforcer.Enforce("zhangsan@qq.com", "/api/v1/setting/permission", "GET")
		fmt.Println(ok)
		fmt.Println(c.Request.URL.RequestURI())
		fmt.Println(c.Request.Method)
		if ok, err := casbin.CasbinObj.Enforcer.Enforce(claims.Email, c.Request.URL.RequestURI(), c.Request.Method); err != nil {
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
