package controller

//
//type LoginResult struct{
//	Name string
//	Token string
//}
//
//func Register(c *gin.Context) {
//	var stu model.Student
//
//	if err:=c.ShouldBind(&stu);err!=nil{
//		c.JSON(http.StatusOK, gin.H{
//			"status": -1,
//			"msg":    err.Error(),
//			"data":   nil,
//		})
//		return
//	}
//
//	service.AddStudent(&stu)
//}
//
//// 定义一个普通controller函数，作为一个验证接口逻辑
//func TestToken(c *gin.Context) {
//	// 上面我们在JWTAuth()中间中将'claims'写入到gin.Context的指针对象中，因此在这里可以将之解析出来
//	claims := c.MustGet("claims").(*middleware.CustomClaims)
//
//	if claims != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status": 0,
//			"msg":    "token有效",
//			"data":   claims,
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"status": -1,
//		"msg":    "token失效",
//		"data":   nil,
//	})
//
//	return
//}
//
//// 登陆
//func Login(c *gin.Context){
//	var stu model.Student
//
//	if err:=c.ShouldBind(&stu);err!=nil{
//		c.JSON(http.StatusOK, gin.H{
//			"status": -1,
//			"msg":    err.Error(),
//			"data":   nil,
//		})
//		return
//	}
//
//	//TODO 插入数据库
//
//
//	generateToken(c,stu)
//
//	return
//}
//
//// token生成器
//func generateToken(c *gin.Context, stu model.Student) {
//	// 构造SignKey: 签名和解签名需要使用一个值
//	j := middleware.NewJWT()
//	// 构造用户claims信息(负荷)
//	claims := middleware.CustomClaims{
//		stu.StuEmail,
//		stu.StuPassword,
//		jwt.StandardClaims{
//			NotBefore: int64(time.Now().Unix() - 1000),
//			// 签名生效时间
//			ExpiresAt: int64(time.Now().Unix() + 3600),
//			// 签名过期时间
//			Issuer:    "bgbiao.top",
//			// 签名颁发者
//			},
//	}
//	// 根据claims生成token对象
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"status": -1,
//			"msg":    err.Error(),
//			"data":   nil,
//		})
//	}
//
//	middleware.Logger().Println(token)
//	// 封装一个响应数据,返回用户名和token
//	data := LoginResult{
//		Name:  stu.StuName,
//		Token: token,
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"status": 0,
//		"msg":    "登陆成功",
//		"data":   data,
//	})
//	return
//}
//
