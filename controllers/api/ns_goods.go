package api

import (
	"errors"
	"fmt"
	"goniushop/models"
	"math"
	"regexp"
	"strings"
	"github.com/astaxie/beego"
)

// NsGoodsController operations for NsGoods
type NsGoodsController struct {
	BaseController
}

// URLMapping ...
func (c *NsGoodsController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("DiscoverList", c.DiscoverList)
}


// DiscoverList ...
// @Title DiscoverList
// @Description get home NsGoods
// @Param	page	query	int	false	"必须是数字"
// @Param	size	query	int	false	"必须是数字"
// @Success 200 {object} models.NsGoods
// @Failure 403
// @router /discoverlist [get]
func (this *NsGoodsController) DiscoverList() {
	page, _ := this.GetInt64("page", 1)
	size, _ := this.GetInt64("size", 10)
	if page <= 0 {
		page = 1
	}
	offset:= (page-1)*size

	list,totalCount,err := models.GetHomeGoodsList(offset,size)
	pageTotal := math.Ceil(float64(totalCount)/float64(size))
	data := map[string]interface{}{"page_total":pageTotal}
	if err != nil {
		this.JsonReturn(1000,err.Error())
	}
	data["list"] = list
	this.JsonReturn(0,"成功",data)
}

// GetOne ...
// @Title Get One
// @Description get NsGoods by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NsGoods
// @Failure 403 :id is empty
// @router /detail [get]
func (c *NsGoodsController) GetOne() {
	id, _ := c.GetInt("id")
	fmt.Println(id)

	v, err := models.GetNsGoodsById(id)
	data := make(map[string]interface{})
	if err != nil {
		c.JsonReturn(1000,err.Error())
	} else {
		re ,_:= regexp.Compile(`<img\s+src=\"(\S+)\"/>`)
		imgUrl := beego.AppConfig.String("ImgUrl")
		imgUrl = imgUrl[0 : len(imgUrl)-1]//不需要最后一个/
		v.Description = re.ReplaceAllString(v.Description,"<img src=\""+imgUrl+"$1\" />")
		data["info"] = v
	}

	images ,err := models.GetGoodsPicture(v.ImgIdArray)
	if err != nil {
		data["images"] = ""
	} else {
		data["images"] = images
	}
	attrs,err := models.GetGoodsAttr(id)
	if err != nil {
		data["attrs"] = ""
	} else {
		data["attrs"] = attrs
	}
	skus,err := models.GetGoodsSkus(id)
	if err != nil {
		data["skus"] = ""
	} else {
		data["skus"] = skus
	}
	c.JsonReturn(0,"成功",data)
}

// GetAll ...
// @Title Get All
// @Description get NsGoods
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.NsGoods
// @Failure 403
// @router / [get]
func (c *NsGoodsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllNsGoods(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}