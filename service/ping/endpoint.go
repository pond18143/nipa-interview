package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) PingGetEndpoint(c *gin.Context) { //GET app/ping
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : pingGet")

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

//รับ INPUT แปลงค่า
func (ep *Endpoint) PingPostEndpoint(c *gin.Context) { //Post app/pingPong
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : pingPost")

	var request inputHeartbeat //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Infof("Body Name : %s", request.Name)
	log.Infof("Body Age  : %d", request.Age)

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

		//return success
	c.JSON(http.StatusOK, result)
	return
}