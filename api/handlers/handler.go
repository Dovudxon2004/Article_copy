package handlers

import(
	"github.com/Dovudxon2004/Api/Article/config"
	"github.com/Dovudxon2004/Api/Article/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	strg storage.StorageI
	cfg config.Config
}

func NewHandler(strg storage.StorageI, cfg config.Config) Handler{
	return Handler{
		strg: strg,
		cfg: cfg,
	}
}

// gets offset param
func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error){
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}


func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}