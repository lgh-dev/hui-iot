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
	"log"
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
	var router = StartServer()
	w := performRequest(router, "POST", PATH_DEVICETYPE_SINGLE_QUERY)
	assert.Equal(t, http.StatusOK, w.Code)
	// 将 JSON 响应转换为 map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	log.Println(w.Body.String())
	assert.NotEqual(t, err, nil)
	assert.NotEqual(t, response["data"], "")
}
