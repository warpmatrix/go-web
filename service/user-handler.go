package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 解析url传递的参数，对于POST则解析响应包的主体（request body）
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		req.ParseForm()
		fname := req.Form["fname"][0]
		lname := req.Form["lname"][0]
		if len(fname) == 0 || len(lname) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		formatter.HTML(w, http.StatusOK, "newUser", struct {
			Fname string
			Lname string
		}{Fname: fname, Lname: lname})
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, nil)
	}
}
