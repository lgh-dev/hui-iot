package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

//单元测试
func TestDeviceModelQuery(t *testing.T) {
	var router = GetServer()
	w := performRequest(router, "GET", "/api/v1/devicemodels")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "ID"), "查询失败!")
}

func BenchmarkDeviceModelQuery(b *testing.B) {
	var router = GetServer()
	b.SetParallelism(4)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			performRequest(router, "GET", "/api/v1/devicemodels")
		}
	})
	//for i:=0;i <b.N ;i++  {
	//	go performRequest(router, "POST", PATH_DEVICETYPE_SINGLE_QUERY)
	//}
}
