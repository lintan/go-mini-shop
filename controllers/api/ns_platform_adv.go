package api

import (
	"goniushop/models"
	"math"
	"strconv"
)

// NsPlatformAdvController operations for NsPlatformAdv
type NsPlatformAdvController struct {
	BaseController
}

// URLMapping ...
func (c *NsPlatformAdvController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}


// GetOne ...
// @Title Get One
// @Description get NsPlatformAdv by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NsPlatformAdv
// @Failure 403 :id is empty
// @router /detail/:id [get]
func (c *NsPlatformAdvController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNsPlatformAdvById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get NsPlatformAdv
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.NsPlatformAdv
// @Failure 403
// @router /list [get]
func (c *NsPlatformAdvController) GetAll() {
	var query = make(map[string]string)
	apId := c.GetString("apId")


	page, _ := c.GetInt64("page", 1)
	size, _ := c.GetInt64("size", 10)
	if page <= 0 {
		page = 1
	}
	offset:= (page-1)*size

	query["ApId"] = apId
	fields := []string{"Id","ApId","AdvTitle","AdvUrl","AdvImage","AdvCode","Background"}
	list, err := models.GetAllNsPlatformAdv(query, fields, []string{"slide_sort","adv_id"}, []string{"desc","desc"}, offset, size)
	data := map[string]interface{}{}
	if err != nil {
		c.JsonReturn(1000,err.Error())
	} else {
		data["list"] = list
	}
	totalCount,err := models.GetAdvCount(query)
	pageTotal := math.Ceil(float64(totalCount)/float64(size))
	data["page_total"] = pageTotal
	if err != nil {
		c.JsonReturn(1000,err.Error())
	}

	c.JsonReturn(0,"成功",data)
}
