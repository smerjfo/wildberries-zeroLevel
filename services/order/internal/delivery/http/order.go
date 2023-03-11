package http

import (
	"github.com/gin-gonic/gin"
	"l0/services/order/internal/delivery/http/order"
	"net/http"
)

func (d *Delivery) ReadOrderByID(c *gin.Context) {
	var id order.ID
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := d.ucOrder.ReadByID(id.Value)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err,
		})
	} else {
		c.HTML(http.StatusOK, "order.html", d.toOrderDAO(response))
	}
}
