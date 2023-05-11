package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetChatReplay(c *gin.Context) {
	jData := map[string]interface{}{
		"ret":     0,
		"errMsg":  "",
		"message": "",
	}

	url := fmt.Sprintf("http://127.0.0.1:8088/api/getChatReplay?question=%s", c.Query("question"))
	rsp, err := http.Get(url)
	if err != nil {
		jData["ret"] = 500
		jData["errMsg"] = err.Error()
		return
	}
	defer rsp.Body.Close()
	resp, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		jData["ret"] = 501
		jData["errMsg"] = "read response fail"
		return
	}

	jData["message"] = string(resp)
	c.JSONP(200, jData)
	return
}
