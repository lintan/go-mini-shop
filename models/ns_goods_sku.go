package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsGoodsSku struct {
	Id                   int     `orm:"column(sku_id);auto" description:"表序号"`
	GoodsId              int     `orm:"column(goods_id)" description:"商品编号"`
	SkuName              string  `orm:"column(sku_name);size(500)" description:"SKU名称"`
	AttrValueItems       string  `orm:"column(attr_value_items);size(255)" description:"属性和属性值 id串 attribute + attribute value 表ID分号分隔"`
	AttrValueItemsFormat string  `orm:"column(attr_value_items_format);size(500)" description:"属性和属性值id串组合json格式"`
	MarketPrice          float64 `orm:"column(market_price);digits(10);decimals(2)" description:"市场价"`
	Price                float64 `orm:"column(price);digits(10);decimals(2)" description:"价格"`
	PromotePrice         float64 `orm:"column(promote_price);digits(10);decimals(2)" description:"促销价格"`
	CostPrice            float64 `orm:"column(cost_price);digits(19);decimals(2)" description:"成本价"`
	Stock                int     `orm:"column(stock)" description:"库存"`
	Picture              int     `orm:"column(picture)" description:"如果是第一个sku编码, 可以加图片"`
	Code                 string  `orm:"column(code);size(255)" description:"商家编码"`
	QRcode               string  `orm:"column(QRcode);size(255)" description:"商品二维码"`
	CreateDate           int     `orm:"column(create_date);null" description:"创建时间"`
	UpdateDate           int     `orm:"column(update_date);null" description:"修改时间"`
}

func (t *NsGoodsSku) TableName() string {
	return "ns_goods_sku"
}

func init() {
	orm.RegisterModel(new(NsGoodsSku))
}

// AddNsGoodsSku insert a new NsGoodsSku into database and returns
// last inserted Id on success.
func AddNsGoodsSku(m *NsGoodsSku) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
func GetGoodsSkus(goodsId int) (skus []map[string]interface{}, err error){
	o := orm.NewOrm()
	//获取图片
	var sk []*NsGoodsSku
	num,err := o.QueryTable(new(NsGoodsSku)).Filter("goods_id", goodsId).All(&sk)

	if num == 0 {
		return nil,err
	}
	for _,v := range sk{
		kv := make(map[string]interface{})

		kv["Id"] = v.Id
		kv["SkuName"] = v.SkuName
		kv["Price"] = v.Price
		kv["Stock"] = v.Stock
		kv["Picture"] = v.Picture
		kv["AttrValueItems"] = v.AttrValueItems
		skus = append(skus, kv)
	}

	return skus,err
}
// GetNsGoodsSkuById retrieves NsGoodsSku by Id. Returns error if
// Id doesn't exist
func GetNsGoodsSkuById(id int) (v *NsGoodsSku, err error) {
	o := orm.NewOrm()
	v = &NsGoodsSku{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsGoodsSku retrieves all NsGoodsSku matches certain condition. Returns empty list if
// no records exist
func GetAllNsGoodsSku(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsGoodsSku))
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

	var l []NsGoodsSku
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

// UpdateNsGoodsSku updates NsGoodsSku by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsGoodsSkuById(m *NsGoodsSku) (err error) {
	o := orm.NewOrm()
	v := NsGoodsSku{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsGoodsSku deletes NsGoodsSku by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsGoodsSku(id int) (err error) {
	o := orm.NewOrm()
	v := NsGoodsSku{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsGoodsSku{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
