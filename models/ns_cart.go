package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsCart struct {
	Id           int     `orm:"column(cart_id);auto" description:"购物车id"`
	BuyerId      int     `orm:"column(buyer_id)" description:"买家id"`
	ShopId       int     `orm:"column(shop_id)" description:"店铺id"`
	ShopName     string  `orm:"column(shop_name);size(100)" description:"店铺名称"`
	GoodsId      int     `orm:"column(goods_id)" description:"商品id"`
	GoodsName    string  `orm:"column(goods_name);size(200)" description:"商品名称"`
	SkuId        int     `orm:"column(sku_id)" description:"商品的skuid"`
	SkuName      string  `orm:"column(sku_name);size(200)" description:"商品的sku名称"`
	Price        float64 `orm:"column(price);digits(10);decimals(2)" description:"商品价格"`
	Num          int16   `orm:"column(num)" description:"购买商品数量"`
	GoodsPicture int     `orm:"column(goods_picture)" description:"商品图片"`
	BlId         int32   `orm:"column(bl_id)" description:"组合套装ID"`
}

func (t *NsCart) TableName() string {
	return "ns_cart"
}

func init() {
	orm.RegisterModel(new(NsCart))
}

// AddNsCart insert a new NsCart into database and returns
// last inserted Id on success.
func AddNsCart(m *NsCart) (id int64, err error) {
	o := orm.NewOrm()
	var cart NsCart
	err1 := o.QueryTable(new(NsCart)).Filter("BuyerId", m.BuyerId).Filter("ShopId", m.ShopId).Filter("GoodsId", m.GoodsId).Filter("SkuId", m.SkuId).Filter("BlId", m.BlId).One(&cart)

	if err1 == orm.ErrNoRows {
		id, err = o.Insert(m)
		return
	}
	_, err2 := o.QueryTable(new(NsCart)).Filter("Id", cart.Id).Update(orm.Params{
		"num": orm.ColValue(orm.ColAdd, m.Num),
	})
	return int64(cart.Id), err2
}

// GetNsCartById retrieves NsCart by Id. Returns error if
// Id doesn't exist
func GetNsCartById(id int) (v *NsCart, err error) {
	o := orm.NewOrm()
	v = &NsCart{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsCart retrieves all NsCart matches certain condition. Returns empty list if
// no records exist
func GetAllNsCart(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsCart))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []NsCart
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNsCart updates NsCart by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsCartById(m *NsCart) (err error) {
	o := orm.NewOrm()
	v := NsCart{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsCart deletes NsCart by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsCart(id int) (err error) {
	o := orm.NewOrm()
	v := NsCart{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsCart{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
