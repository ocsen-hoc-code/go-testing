package user_test

import (
	"encoding/json"
	"testing"

	"github.com/ocsen-hoc-code/go-auj/models/service"
	util_test "github.com/ocsen-hoc-code/go-auj/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	serv := util_test.InstallTest()
	defer func(s *service.Service) {
		serv.Config.Database.Close()
	}(serv)

	t.Run("TestLogin", func(t *testing.T) {
		values := map[string]string{"username": util_test.USERNAME, "password": util_test.PASSWORD}
		jsonData, _ := json.Marshal(values)
		status, body := util_test.CreateRequest(serv.Server, util_test.POST, "/login", "", jsonData)
		assert.Equal(t, 200, status, "Expect 200")
		assert.NotEmpty(t, body["data"], "Login success return token")
	})

	t.Run("TestLoginFail", func(t *testing.T) {
		values := map[string]string{"username": util_test.USERNAME, "password": util_test.PASSWORD + "fail"}
		jsonData, _ := json.Marshal(values)
		status, body := util_test.CreateRequest(serv.Server, util_test.POST, "/login", "", jsonData)
		assert.Equal(t, 400, status, "Expect 400")
		assert.Empty(t, body["data"], "Expect body is emty")
	})
}
