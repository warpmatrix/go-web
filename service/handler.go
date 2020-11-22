package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", nil)
	}
}

func getDataHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, uList)
	}
}

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// 解析url传递的参数，对于POST则解析响应包的主体（request body）
		// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
		req.ParseForm()
		if !isValid(req.Form) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		newUser := parseUser(req.Form)
		uList = append(uList, newUser)
		formatter.HTML(w, http.StatusOK, "newUser", newUser)
	}
}

func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "userInfo", nil)
	}
}

// notImplemented replies to the request with an HTTP 501 Not Implemented.
func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", http.StatusNotImplemented)
}

// notImplementedHandler returns a simple request handler
// that replies to each request with a ``501 Not Implemented'' reply.
func notImplementedHandler() http.HandlerFunc { return http.HandlerFunc(notImplemented) }
