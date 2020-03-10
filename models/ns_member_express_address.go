package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsMemberExpressAddress struct {
	Id        int    `orm:"column(id);auto"`
	Uid       int    `orm:"column(uid)" description:"会员基本资料表ID"`
	Consigner string `orm:"column(consigner);size(255)" description:"收件人"`
	Mobile    string `orm:"column(mobile);size(11)" description:"手机"`
	Phone     string `orm:"column(phone);size(20)" description:"固定电话"`
	Province  int    `orm:"column(province)" description:"省"`
	City      int    `orm:"column(city)" description:"市"`
	District  int    `orm:"column(district)" description:"区县"`
	Address   string `orm:"column(address);size(255)" description:"详细地址"`
	ZipCode   string `orm:"column(zip_code);size(6)" description:"邮编"`
	Alias     string `orm:"column(alias);size(50)" description:"地址别名"`
	IsDefault int    `orm:"column(is_default)" description:"默认收货地址"`
	FullAddress string `orm:"-"`
}

func (t *NsMemberExpressAddress) TableName() string {
	return "ns_member_express_address"
}

func init() {
	orm.RegisterModel(new(NsMemberExpressAddress))
}

// AddNsMemberExpressAddress insert a new NsMemberExpressAddress into database and returns
// last inserted Id on success.
func AddNsMemberExpressAddress(m *NsMemberExpressAddress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func GetFullAddress(m *NsMemberExpressAddress) string{
	provinceName := GetProvinceNameById(m.Province)
	cityName := GetCityNameById(m.Province)
	districtName := GetDistrictNameById(m.Province)
	return provinceName+cityName+districtName+m.Address
}
func GetMemberAddressByUid(uid int)(address []NsMemberExpressAddress, err error){
	o := orm.NewOrm()
	//var address []NsMemberExpressAddress
	_,err = o.QueryTable(new(NsMemberExpressAddress)).Filter("uid",uid).All(&address)
	//fmt.Println(address)
	return
}

// GetNsMemberExpressAddressById retrieves NsMemberExpressAddress by Id. Returns error if
// Id doesn't exist
func GetNsMemberExpressAddressById(id int) (v *NsMemberExpressAddress, err error) {
	o := orm.NewOrm()
	v = &NsMemberExpressAddress{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsMemberExpressAddress retrieves all NsMemberExpressAddress matches certain condition. Returns empty list if
// no records exist
func GetAllNsMemberExpressAddress(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsMemberExpressAddress))
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

	var l []NsMemberExpressAddress
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

// UpdateNsMemberExpressAddress updates NsMemberExpressAddress by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsMemberExpressAddressById(m *NsMemberExpressAddress) (err error) {
	o := orm.NewOrm()
	v := NsMemberExpressAddress{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsMemberExpressAddress deletes NsMemberExpressAddress by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsMemberExpressAddress(id int) (err error) {
	o := orm.NewOrm()
	v := NsMemberExpressAddress{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsMemberExpressAddress{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
