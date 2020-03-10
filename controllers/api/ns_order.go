package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"goniushop/models"
	"strconv"
	"strings"

)

// NsOrderController operations for NsOrder
type NsOrderController struct {
	BaseController
}

// URLMapping ...
func (c *NsOrderController) URLMapping() {
	c.Mapping("SaveOrder", c.SaveOrder)
	c.Mapping("PreOrder", c.PreOrder)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create NsOrder
// @Param	body		body 	models.NsOrder	true		"body for NsOrder content"
// @Success 201 {int} models.NsOrder
// @Failure 403 body is empty
// @router /saveOrder [post]
func (c *NsOrderController) SaveOrder() {
	//var v models.NsOrder
	type FormData struct {
		BuyerMessage string  `form:"buyerMessage"`
		FormId string `form:"formId"`
		Sign string `form:"sign"`
		Time string `form:"time"`
		GoodsJsonStr string `form:"goodsJsonStr"`
		ReceiverInfoId int `form:"receiverInfoId"`
	}
	type GoodsList struct {
		GoodsId int `json:goodsId`
		ShopId int `json:shopId`
		GoodsName string `json:goodsName`
		GoodsPicture int `json:goodsPicture`
		SkuId int `json:skuId`
		SkuName string `json:skuName`
		Price float64 `json:price`
		Num int `json:num`
	}
	var form FormData
	var gl []GoodsList
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err == nil {
		orderNo := models.GetUniOrderNo()
		outTradeNo := models.GetUniOutTradeNo()
		fmt.Println("===",orderNo,"===",outTradeNo)
		if err := json.Unmarshal([]byte(form.GoodsJsonStr), &gl); err == nil {
			fmt.Println(gl[0])
		}else{
			fmt.Println(err.Error())
		}


		/*if _, err := models.AddNsOrder(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}*/
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// PreOrder ...
// @Title PreOrder
// @Description get NsOrder by id
// @Success 200 {object} models.NsOrder
// @Failure 403 :id is empty
// @router /preOrder [get]
func (c *NsOrderController) PreOrder() {
	data := make(map[string]interface{})
	addressId,_ := c.GetInt("addressId",0)

	data["hasDefaultAddress"] = 0
	data["address"] = ""
	if addressId != 0 {
		v,_ := models.GetNsMemberExpressAddressById(addressId)
		v.FullAddress = models.GetFullAddress(v)
		data["address"] = v
		data["hasDefaultAddress"] = v.IsDefault
	}else{
		l,err:=models.GetMemberAddressByUid(2)
		if err != nil {
			fmt.Println(err.Error())
		}

		for _,v := range l {
			if v.IsDefault == 1 {
				data["hasDefaultAddress"] = 1
				v.FullAddress = models.GetFullAddress(&v)
				data["address"] = v
			}
		}
	}
	data["ShippingFee"] = 1

	c.JsonResult(0,"成功",data)
}
// GetOne ...
// @Title Get One
// @Description get NsOrder by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.NsOrder
// @Failure 403 :id is empty
// @router /preOrder/:id [get]
func (c *NsOrderController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetNsOrderById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get NsOrder
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.NsOrder
// @Failure 403
// @router / [get]
func (c *NsOrderController) GetAll() {
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

	l, err := models.GetAllNsOrder(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the NsOrder
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.NsOrder	true		"body for NsOrder content"
// @Success 200 {object} models.NsOrder
// @Failure 403 :id is not int
// @router /:id [put]
func (c *NsOrderController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.NsOrder{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateNsOrderById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the NsOrder
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *NsOrderController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteNsOrder(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
