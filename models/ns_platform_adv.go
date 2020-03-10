package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type NsPlatformAdv struct {
	Id         int    `orm:"column(adv_id);auto" description:"广告自增标识编号"`
	ApId       uint32 `orm:"column(ap_id)" description:"广告位id"`
	AdvTitle   string `orm:"column(adv_title);size(255)" description:"广告内容描述"`
	AdvUrl     string `orm:"column(adv_url);null"`
	AdvImage   string `orm:"column(adv_image);size(1000)" description:"广告内容图片"`
	SlideSort  int    `orm:"column(slide_sort);null"`
	ClickNum   uint   `orm:"column(click_num)" description:"广告点击率"`
	Background string `orm:"column(background);size(255)" description:"背景色"`
	AdvCode    string `orm:"column(adv_code)" description:"广告代码"`
}

func (t *NsPlatformAdv) TableName() string {
	return "ns_platform_adv"
}

func init() {
	orm.RegisterModel(new(NsPlatformAdv))
}

// AddNsPlatformAdv insert a new NsPlatformAdv into database and returns
// last inserted Id on success.
func AddNsPlatformAdv(m *NsPlatformAdv) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsPlatformAdvById retrieves NsPlatformAdv by Id. Returns error if
// Id doesn't exist
func GetNsPlatformAdvById(id int) (v *NsPlatformAdv, err error) {
	o := orm.NewOrm()
	v = &NsPlatformAdv{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsPlatformAdv retrieves all NsPlatformAdv matches certain condition. Returns empty list if
// no records exist
func GetAllNsPlatformAdv(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsPlatformAdv))
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

	var l []NsPlatformAdv
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				imgUrl := beego.AppConfig.String("ImgUrl")
				v.AdvImage = fmt.Sprintf("%s%s",imgUrl,v.AdvImage)
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

func GetAdvCount(query map[string]string) (count int64 ,  err error){
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsPlatformAdv))
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	count,err  = qs.Count()

	return count , err
}

// UpdateNsPlatformAdv updates NsPlatformAdv by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsPlatformAdvById(m *NsPlatformAdv) (err error) {
	o := orm.NewOrm()
	v := NsPlatformAdv{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsPlatformAdv deletes NsPlatformAdv by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsPlatformAdv(id int) (err error) {
	o := orm.NewOrm()
	v := NsPlatformAdv{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsPlatformAdv{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
