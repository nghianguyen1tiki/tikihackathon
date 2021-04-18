package tikingon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/server"
)

type handler struct {
}

var _ server.Handler = (*handler)(nil)

func NewHandler() server.Handler {
	return &handler{}
}

func (h *handler) Register(router gin.IRouter) {
	router.GET("", h.searchTikiNgon)
}

type Entity struct {
	Name string `json:"name"`
	Price int `json:"price"`
	DiscountRate int `json:"discount_rate"`
	ThumbnailUrl string `json:"thumbnail_url"`
	ThumbnailWidth int `json:"thumbnail_width"`
	ThumbnailHeight int `json:"thumbnail_height"`
}

type Res struct {
	Data []*Entity `json:"data"`
}

func (h *handler) searchTikiNgon(c *gin.Context) {
	ingredientName := c.DefaultQuery("q", "")
	category := c.DefaultQuery("category", "44792")
	limit := c.DefaultQuery("limit", "10")

	ft := "https://tiki.vn/api/v2/products?limit=%v&q=%v&category=%v"
	url := fmt.Sprintf(ft,limit,url.QueryEscape(ingredientName),category)
	method := "GET"

	client := &http.Client{
		Timeout:       5*time.Second,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	result := &Res{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}
