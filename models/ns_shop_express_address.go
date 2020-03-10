package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsShopExpressAddress struct {
	Id          int    `orm:"column(express_address_id);auto" description:"物流地址id"`
	ShopId      int    `orm:"column(shop_id)" description:"商铺id"`
	Contact     string `orm:"column(contact);size(100)" description:"联系人"`
	Mobile      string `orm:"column(mobile);size(50)" description:"手机"`
	Phone       string `orm:"column(phone);size(50)" description:"电话"`
	CompanyName string `orm:"column(company_name);size(100)" description:"公司名称"`
	Province    int16  `orm:"column(province)" description:"所在地省"`
	City        int16  `orm:"column(city)" description:"所在地市"`
	District    int16  `orm:"column(district)" description:"所在地区县"`
	Zipcode     string `orm:"column(zipcode);size(6)" description:"邮编"`
	Address     string `orm:"column(address);size(100)" description:"详细地址"`
	IsConsigner int8   `orm:"column(is_consigner)" description:"发货地址标记"`
	IsReceiver  int8   `orm:"column(is_receiver)" description:"收货地址标记"`
	CreateDate  int    `orm:"column(create_date);null" description:"创建日期"`
	ModifyDate  int    `orm:"column(modify_date);null" description:"修改日期"`
}

func (t *NsShopExpressAddress) TableName() string {
	return "ns_shop_express_address"
}

func init() {
	orm.RegisterModel(new(NsShopExpressAddress))
}

// AddNsShopExpressAddress insert a new NsShopExpressAddress into database and returns
// last inserted Id on success.
func AddNsShopExpressAddress(m *NsShopExpressAddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsShopExpressAddressById retrieves NsShopExpressAddress by Id. Returns error if
// Id doesn't exist
func GetNsShopExpressAddressById(id int) (v *NsShopExpressAddress, err error) {
	o := orm.NewOrm()
	v = &NsShopExpressAddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsShopExpressAddress retrieves all NsShopExpressAddress matches certain condition. Returns empty list if
// no records exist
func GetAllNsShopExpressAddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsShopExpressAddress))
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

	var l []NsShopExpressAddress
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

// UpdateNsShopExpressAddress updates NsShopExpressAddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsShopExpressAddressById(m *NsShopExpressAddress) (err error) {
	o := orm.NewOrm()
	v := NsShopExpressAddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsShopExpressAddress deletes NsShopExpressAddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsShopExpressAddress(id int) (err error) {
	o := orm.NewOrm()
	v := NsShopExpressAddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsShopExpressAddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
