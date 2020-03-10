package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsOrderGoodsExpress struct {
	Id                int    `orm:"column(id);auto"`
	OrderId           int    `orm:"column(order_id)" description:"订单id"`
	OrderGoodsIdArray string `orm:"column(order_goods_id_array);size(255)" description:"订单项商品组合列表"`
	ExpressName       string `orm:"column(express_name);size(50)" description:"包裹名称  （包裹- 1 包裹 - 2）"`
	ShippingType      int8   `orm:"column(shipping_type)" description:"发货方式1 需要物流 0无需物流"`
	ExpressCompanyId  int    `orm:"column(express_company_id)" description:"快递公司id"`
	ExpressCompany    string `orm:"column(express_company);size(255)" description:"物流公司名称"`
	ExpressNo         string `orm:"column(express_no);size(50)" description:"运单编号"`
	Uid               int    `orm:"column(uid)" description:"用户id"`
	UserName          string `orm:"column(user_name);size(50)" description:"用户名"`
	Memo              string `orm:"column(memo);size(255)" description:"备注"`
	ShippingTime      int    `orm:"column(shipping_time);null" description:"发货时间"`
}

func (t *NsOrderGoodsExpress) TableName() string {
	return "ns_order_goods_express"
}

func init() {
	orm.RegisterModel(new(NsOrderGoodsExpress))
}

// AddNsOrderGoodsExpress insert a new NsOrderGoodsExpress into database and returns
// last inserted Id on success.
func AddNsOrderGoodsExpress(m *NsOrderGoodsExpress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsOrderGoodsExpressById retrieves NsOrderGoodsExpress by Id. Returns error if
// Id doesn't exist
func GetNsOrderGoodsExpressById(id int) (v *NsOrderGoodsExpress, err error) {
	o := orm.NewOrm()
	v = &NsOrderGoodsExpress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsOrderGoodsExpress retrieves all NsOrderGoodsExpress matches certain condition. Returns empty list if
// no records exist
func GetAllNsOrderGoodsExpress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsOrderGoodsExpress))
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

	var l []NsOrderGoodsExpress
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

// UpdateNsOrderGoodsExpress updates NsOrderGoodsExpress by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsOrderGoodsExpressById(m *NsOrderGoodsExpress) (err error) {
	o := orm.NewOrm()
	v := NsOrderGoodsExpress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsOrderGoodsExpress deletes NsOrderGoodsExpress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsOrderGoodsExpress(id int) (err error) {
	o := orm.NewOrm()
	v := NsOrderGoodsExpress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsOrderGoodsExpress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
