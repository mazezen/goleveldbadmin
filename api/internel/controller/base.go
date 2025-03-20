package controller

import "net/http"

func cors(w http.ResponseWriter, r *http.Request) {
	// 允许所有来源访问，你也可以指定特定的来源
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 允许特定的请求方法
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// 允许特定的请求头
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 预检请求的缓存时间（单位秒）
	w.Header().Set("Access-Control-Max-Age", "3600")
	// 如果是预检请求，直接返回成功状态码
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}
