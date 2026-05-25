package controllers

import (
	"errors"
	"net/http"

	"github.com/bhcoder23/gin-layout/internal/interfaces"
	"github.com/bhcoder23/gin-layout/internal/models"
	"github.com/bhcoder23/gin-layout/internal/services"
	"github.com/bhcoder23/gin-layout/internal/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func GoodsPageList(c *gin.Context) {
	var filter interfaces.IFilter = &models.GoodsFilter{}
	err := c.ShouldBindQuery(&filter)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}

	searchKey := c.DefaultQuery(`searchKey`, ``)
	filter.SetSearchKey(searchKey)
	page := c.DefaultQuery(`page`, `1`)
	filter.SetPage(cast.ToInt64(page))
	pageSize := c.DefaultQuery(`pageSize`, `20`)
	filter.SetPageSize(cast.ToInt64(pageSize))
	if 0 >= filter.GetPage() {
		filter.SetPage(1)
	}
	if 0 >= filter.GetPageSize() {
		filter.SetPageSize(20)
	}
	if filter.GetPageSize() > 1000 {
		response.Error(`page size cannot exceed 1000`, http.StatusBadRequest, c)
		return
	}
	// nType := c.DefaultQuery(`type`, `0`)
	// filter.SetType(nType)
	result, err := services.Goods.PageList(c, &filter)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(result, c)
}

func GoodsOne(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`missing id`, http.StatusBadRequest, c)
		return
	}

	one, err := services.Goods.One(c, cast.ToUint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(one, c)
}

func GoodsAdd(c *gin.Context) {
	model := &models.Goods{}
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	newID, err := services.Goods.Add(c, model)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(newID, c)
}

func GoodsUpdate(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`missing id`, http.StatusBadRequest, c)
		return
	}
	model := make(map[string]any)
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	delete(model, `updated_at`)
	updated, err := services.Goods.Update(c, model, cast.ToUint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

func GoodsDel(c *gin.Context) {
	id := c.Param(`id`)
	if id == `` {
		response.Error(`missing id`, http.StatusBadRequest, c)
		return
	}
	deleted, err := services.Goods.Delete(c, cast.ToUint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(err.Error(), http.StatusNotFound, c)
			return
		}
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}
