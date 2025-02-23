package log_handelr

import (
	"net/http"
	"strconv"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_monitor"
	"vue-admin/web_server/server/handler/handler_common"

	"github.com/flywithbug/log4go"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func getLogListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponse, aRes)
		c.JSON(http.StatusOK, aRes)
	}()
	size, _ := strconv.Atoi(c.Query("size"))
	page, _ := strconv.Atoi(c.Query("page"))
	userId := c.Query("user_id")
	requestId := c.Query("requestId")
	resCode, _ := strconv.Atoi(c.Query("response_code"))

	flag := c.Query("flag")
	if size == 0 {
		size = 10
	}
	if page != 0 {
		page--
	}
	query := bson.M{}
	if len(userId) > 0 {
		uID, _ := strconv.Atoi(c.Query("user_id"))
		query["user_id"] = uID
	}
	if len(requestId) > 0 {
		query["request_id"] = requestId
	}
	if len(flag) > 0 {
		query["flag"] = flag
	}
	if resCode > 0 {
		query["response_code"] = resCode
	}
	var l = model_monitor.Log{}
	totalCount, _ := l.TotalCount(query, nil)
	results, err := l.FindPageFilter(page, size, query, nil, "-_id")
	if err != nil {
		log4go.Error(handler_common.RequestId(c) + err.Error())
		aRes.SetErrorInfo(http.StatusInternalServerError, "list find error"+err.Error())
		return
	}
	aRes.AddResponseInfo("list", results)
	aRes.AddResponseInfo("total", totalCount)
}
