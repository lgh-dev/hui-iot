package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"hui-iot/iot-server/dto"
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
func TestFindAllDeviceModels(t *testing.T) {
	var router = GetServer()
	w := performRequest(router, "GET", "/api/v1/devicemodels")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "ID"), "查询失败!")
}

func TestFindDeviceModelByIds(t *testing.T) {
	var router = GetServer()
	w := performRequest(router, "GET", "/api/v1/devicemodel/car")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.Contains(w.Body.String(), "ID"), "查询失败!")
}

func TestAddDevice(t *testing.T) {
	var router = GetServer()
	deviceDTO := dto.DeviceDTO{Uid: "1", DeviceModelID: "car", Name: "lgh"}
	bodyBytes, _ := json.Marshal(deviceDTO)
	req, _ := http.NewRequest("POST", "/api/v1/device", bytes.NewReader(bodyBytes))
	req.Header.Add("Authorization", "Bearer 111")
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func BenchmarkFindAllDeviceModels(b *testing.B) {
	var router = GetServer()
	b.SetParallelism(4)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			performRequest(router, "GET", "/api/v1/devicemodels")
		}
	})
}

func BenchmarkFindDeviceModelByIds(b *testing.B) {
	var router = GetServer()
	b.SetParallelism(4)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			performRequest(router, "GET", "/api/v1/devicemodel/car")
		}
	})
}
