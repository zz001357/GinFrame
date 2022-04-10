/**
* @Author: ZhangZe
* @Date: 2022/4/10 18:44
 */

package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "hello World!111111")
}
