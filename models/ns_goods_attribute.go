package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsGoodsAttribute struct {
	Id            int    `orm:"column(attr_id);auto"`
	GoodsId       int    `orm:"column(goods_id)" description:"商品ID"`
	ShopId        int    `orm:"column(shop_id)" description:"店铺ID"`
	AttrValueId   int    `orm:"column(attr_value_id)" description:"属性值id"`
	AttrValue     string `orm:"column(attr_value);size(255)" description:"属性值名称"`
	AttrValueName string `orm:"column(attr_value_name);size(255)" description:"属性值对应数据值"`
	Sort          int    `orm:"column(sort)" description:"排序"`
	CreateTime    int    `orm:"column(create_time);null" description:"创建时间"`
}

func (t *NsGoodsAttribute) TableName() string {
	return "ns_goods_attribute"
}

func init() {
	orm.RegisterModel(new(NsGoodsAttribute))
}

// AddNsGoodsAttribute insert a new NsGoodsAttribute into database and returns
// last inserted Id on success.
func AddNsGoodsAttribute(m *NsGoodsAttribute) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsGoodsAttributeById retrieves NsGoodsAttribute by Id. Returns error if
// Id doesn't exist
func GetNsGoodsAttributeById(id int) (v *NsGoodsAttribute, err error) {
	o := orm.NewOrm()
	v = &NsGoodsAttribute{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetGoodsAttr(goodsId int) (attrs []map[string]string, err error){
	o := orm.NewOrm()
	//获取图片
	var at []*NsGoodsAttribute
	num,err := o.QueryTable(new(NsGoodsAttribute)).Filter("goods_id", goodsId).All(&at)

	if num == 0 {
		return nil,err
	}
	for _,v := range at{
		kv := make(map[string]string)
		kv["AttrValue"] = v.AttrValue
		kv["AttrValueName"] = v.AttrValueName
		attrs = append(attrs, kv)
	}

	return attrs,err
}

// GetAllNsGoodsAttribute retrieves all NsGoodsAttribute matches certain condition. Returns empty list if
// no records exist
func GetAllNsGoodsAttribute(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsGoodsAttribute))
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

	var l []NsGoodsAttribute
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

// UpdateNsGoodsAttribute updates NsGoodsAttribute by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsGoodsAttributeById(m *NsGoodsAttribute) (err error) {
	o := orm.NewOrm()
	v := NsGoodsAttribute{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsGoodsAttribute deletes NsGoodsAttribute by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsGoodsAttribute(id int) (err error) {
	o := orm.NewOrm()
	v := NsGoodsAttribute{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsGoodsAttribute{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
