package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"net/http"
	"net/http/httptest"
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
func TestDeviceTypeQuery(t *testing.T) {
	var router = GetServer()
	w := performRequest(router, "GET", "/api/svc-iot-device/v1/devicetype/")

	assert.Equal(t, http.StatusOK, w.Code)
	// 将 JSON 响应转换为 map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	assert.NotEqual(t, err, nil)
	//assert.NotEqual(t, response["data"], "")
}

func BenchmarkDeviceTypeQuery(b *testing.B) {
	var router = GetServer()
	b.SetParallelism(4)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			performRequest(router, "GET", "/api/svc-iot-device/v1/devicetype/")
		}
	})
	//for i:=0;i <b.N ;i++  {
	//	go performRequest(router, "POST", PATH_DEVICETYPE_SINGLE_QUERY)
	//}
}
