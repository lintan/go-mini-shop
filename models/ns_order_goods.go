package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type NsOrderGoods struct {
	Id                    int     `orm:"column(order_goods_id);auto" description:"订单项ID"`
	OrderId               int     `orm:"column(order_id)" description:"订单ID"`
	GoodsId               int     `orm:"column(goods_id)" description:"商品ID"`
	GoodsName             string  `orm:"column(goods_name);size(100)" description:"商品名称"`
	SkuId                 int     `orm:"column(sku_id)" description:"skuID"`
	SkuName               string  `orm:"column(sku_name);size(50)" description:"sku名称"`
	Price                 float64 `orm:"column(price);digits(19);decimals(2)" description:"商品价格"`
	CostPrice             float64 `orm:"column(cost_price);digits(19);decimals(2)" description:"商品成本价"`
	Num                   string  `orm:"column(num);size(255)" description:"购买数量"`
	AdjustMoney           float64 `orm:"column(adjust_money);digits(10);decimals(2)" description:"调整金额"`
	GoodsMoney            float64 `orm:"column(goods_money);digits(10);decimals(2)" description:"商品总价"`
	GoodsPicture          int     `orm:"column(goods_picture)" description:"商品图片"`
	ShopId                int     `orm:"column(shop_id)" description:"店铺ID"`
	BuyerId               int     `orm:"column(buyer_id)" description:"购买人ID"`
	PointExchangeType     int     `orm:"column(point_exchange_type)" description:"积分兑换类型0.非积分兑换1.积分兑换"`
	GoodsType             string  `orm:"column(goods_type);size(255)" description:"商品类型"`
	PromotionId           int     `orm:"column(promotion_id)" description:"促销ID"`
	PromotionTypeId       int     `orm:"column(promotion_type_id)" description:"促销类型"`
	OrderType             int     `orm:"column(order_type)" description:"订单类型"`
	OrderStatus           int     `orm:"column(order_status)" description:"订单状态"`
	GivePoint             int     `orm:"column(give_point)" description:"积分数量"`
	ShippingStatus        int     `orm:"column(shipping_status)" description:"物流状态"`
	RefundType            int     `orm:"column(refund_type)" description:"退款方式"`
	RefundRequireMoney    float64 `orm:"column(refund_require_money);digits(10);decimals(2)" description:"退款金额"`
	RefundReason          string  `orm:"column(refund_reason);size(255)" description:"退款原因"`
	RefundShippingCode    string  `orm:"column(refund_shipping_code);size(255)" description:"退款物流单号"`
	RefundShippingCompany string  `orm:"column(refund_shipping_company);size(255)" description:"退款物流公司名称"`
	RefundRealMoney       float64 `orm:"column(refund_real_money);digits(10);decimals(2)" description:"实际退款金额"`
	RefundStatus          int     `orm:"column(refund_status)" description:"退款状态"`
	Memo                  string  `orm:"column(memo);size(255)" description:"备注"`
	IsEvaluate            int16   `orm:"column(is_evaluate)" description:"是否评价 0为未评价 1为已评价 2为已追评"`
	RefundTime            int     `orm:"column(refund_time);null" description:"退款时间"`
	RefundBalanceMoney    float64 `orm:"column(refund_balance_money);digits(10);decimals(2)" description:"订单退款余额"`
	TmpExpressCompany     string  `orm:"column(tmp_express_company);size(255)" description:"批量打印时添加的临时物流公司"`
	TmpExpressCompanyId   int     `orm:"column(tmp_express_company_id)" description:"批量打印时添加的临时物流公司id"`
	TmpExpressNo          string  `orm:"column(tmp_express_no);size(50)" description:"批量打印时添加的临时订单编号"`
	GiftFlag              int     `orm:"column(gift_flag);null" description:"赠品标识，0:不是赠品，大于0：赠品id"`
}

func (t *NsOrderGoods) TableName() string {
	return "ns_order_goods"
}

func init() {
	orm.RegisterModel(new(NsOrderGoods))
}

// AddNsOrderGoods insert a new NsOrderGoods into database and returns
// last inserted Id on success.
func AddNsOrderGoods(m *NsOrderGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsOrderGoodsById retrieves NsOrderGoods by Id. Returns error if
// Id doesn't exist
func GetNsOrderGoodsById(id int) (v *NsOrderGoods, err error) {
	o := orm.NewOrm()
	v = &NsOrderGoods{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsOrderGoods retrieves all NsOrderGoods matches certain condition. Returns empty list if
// no records exist
func GetAllNsOrderGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsOrderGoods))
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

	var l []NsOrderGoods
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

// UpdateNsOrderGoods updates NsOrderGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsOrderGoodsById(m *NsOrderGoods) (err error) {
	o := orm.NewOrm()
	v := NsOrderGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsOrderGoods deletes NsOrderGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsOrderGoods(id int) (err error) {
	o := orm.NewOrm()
	v := NsOrderGoods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsOrderGoods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
