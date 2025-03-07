package monitor_handler

import (
	"net/http"
	"time"
	"vue-admin/web_server/common"
	"vue-admin/web_server/model"
	"vue-admin/web_server/model/model_monitor"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
)

const (
	MonitorVisit = "visit"
)

type responseVisit struct {
	DayVisit   int `json:"dayVisit"`
	TotalVisit int `json:"totalVisit"`
	DayApi     int `json:"dayApi"`
	TotalApi   int `json:"totalApi"`

	DayIP   int `json:"dayIp"`
	TotalIp int `json:"totalIp"`
}

func visitHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponse, aRes)
		c.JSON(http.StatusOK, aRes)
	}()
	query := bson.M{}
	monitorCount := model_monitor.MonitorCount{}

	resVisit := responseVisit{}
	timeF := time.Now().Format(model_monitor.TimeLayout)
	timeF = timeF[:10]
	vApi := model_monitor.VisitApi{}
	vUId := model_monitor.VisitUId{}

	query = bson.M{"time_date": bson.M{"$regex": timeF, "$options": "i"}}
	resVisit.DayApi, _ = vApi.TotalSumCount(query)
	monitorCount, _ = monitorCount.FindOne(bson.M{"monitor": MonitorVisit, "time_date": timeF}) //日访问
	resVisit.DayVisit = monitorCount.Count
	resVisit.DayIP, _ = vUId.TotalCount(query, nil)

	query = bson.M{"time_date": bson.M{"$regex": "", "$options": "i"}}
	resVisit.TotalApi, _ = vApi.TotalSumCount(query)
	resVisit.TotalVisit, _ = monitorCount.TotalSumCount(query) //总访问

	resVisit.TotalIp, _ = vUId.TotalCount(nil, nil)

	aRes.AddResponseInfo("visit", resVisit)
}

func visitCountHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponse, aRes)
		c.JSON(http.StatusOK, aRes)
	}()

	mon := model_monitor.MonitorCount{}
	mon.Monitor = MonitorVisit
	mon.IncrementMonitorCount()
	aRes.SetSuccess()
}

func chartListHandler(c *gin.Context) {
	aRes := model.NewResponse()
	defer func() {
		c.Set(common.KeyContextResponse, aRes)
		c.JSON(http.StatusOK, aRes)
	}()

}
