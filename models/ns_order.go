package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/rs/xid"
)

type NsOrder struct {
	Id          int    `orm:"column(order_id);auto" description:"订单id"`
	OrderNo     string `orm:"column(order_no);size(255);null" description:"订单编号"`
	OutTradeNo  string `orm:"column(out_trade_no);size(100)" description:"外部交易号"`
	OrderType   int8   `orm:"column(order_type)" description:"订单类型"`
	PaymentType int8   `orm:"column(payment_type)" description:"支付类型。取值范围：WEIXIN (微信自有支付)WEIXIN_DAIXIAO (微信代销支付)ALIPAY (支付宝支付)"`
	ShippingType        int8    `orm:"column(shipping_type)" description:"订单配送方式"`
	OrderFrom           string  `orm:"column(order_from);size(255)" description:"订单来源"`
	BuyerId             int     `orm:"column(buyer_id)" description:"买家id"`
	UserName            string  `orm:"column(user_name);size(50)" description:"买家会员名称"`
	BuyerIp             string  `orm:"column(buyer_ip);size(20)" description:"买家ip"`
	BuyerMessage        string  `orm:"column(buyer_message);size(255)" description:"买家附言"`
	BuyerInvoice        string  `orm:"column(buyer_invoice);size(255)" description:"买家发票信息"`
	ReceiverMobile      string  `orm:"column(receiver_mobile);size(11)" description:"收货人的手机号码"`
	ReceiverProvince    int     `orm:"column(receiver_province)" description:"收货人所在省"`
	ReceiverCity        int     `orm:"column(receiver_city)" description:"收货人所在城市"`
	ReceiverDistrict    int     `orm:"column(receiver_district)" description:"收货人所在街道"`
	ReceiverAddress     string  `orm:"column(receiver_address);size(255)" description:"收货人详细地址"`
	ReceiverZip         string  `orm:"column(receiver_zip);size(6)" description:"收货人邮编"`
	ReceiverName        string  `orm:"column(receiver_name);size(50)" description:"收货人姓名"`
	ShopId              int     `orm:"column(shop_id)" description:"卖家店铺id"`
	ShopName            string  `orm:"column(shop_name);size(100)" description:"卖家店铺名称"`
	SellerStar          int8    `orm:"column(seller_star)" description:"卖家对订单的标注星标"`
	SellerMemo          string  `orm:"column(seller_memo);size(255)" description:"卖家对订单的备注"`
	ConsignTimeAdjust   int     `orm:"column(consign_time_adjust)" description:"卖家延迟发货时间"`
	GoodsMoney          float64 `orm:"column(goods_money);digits(19);decimals(2)" description:"商品总价"`
	OrderMoney          float64 `orm:"column(order_money);digits(10);decimals(2)" description:"订单总价"`
	Point               int     `orm:"column(point)" description:"订单消耗积分"`
	PointMoney          float64 `orm:"column(point_money);digits(10);decimals(2)" description:"订单消耗积分抵多少钱"`
	CouponMoney         float64 `orm:"column(coupon_money);digits(10);decimals(2)" description:"订单代金券支付金额"`
	CouponId            int     `orm:"column(coupon_id)" description:"订单代金券id"`
	UserMoney           float64 `orm:"column(user_money);digits(10);decimals(2)" description:"订单余额支付金额"`
	UserPlatformMoney   float64 `orm:"column(user_platform_money);digits(10);decimals(2)" description:"用户平台余额支付"`
	PromotionMoney      float64 `orm:"column(promotion_money);digits(10);decimals(2)" description:"订单优惠活动金额"`
	ShippingMoney       float64 `orm:"column(shipping_money);digits(10);decimals(2)" description:"订单运费"`
	PayMoney            float64 `orm:"column(pay_money);digits(10);decimals(2)" description:"订单实付金额"`
	RefundMoney         float64 `orm:"column(refund_money);digits(10);decimals(2)" description:"订单退款金额"`
	CoinMoney           float64 `orm:"column(coin_money);digits(10);decimals(2)" description:"购物币金额"`
	GivePoint           int     `orm:"column(give_point)" description:"订单赠送积分"`
	GiveCoin            float64 `orm:"column(give_coin);digits(10);decimals(2)" description:"订单成功之后返购物币"`
	OrderStatus         int8    `orm:"column(order_status)" description:"订单状态"`
	PayStatus           int8    `orm:"column(pay_status)" description:"订单付款状态"`
	ShippingStatus      int8    `orm:"column(shipping_status)" description:"订单配送状态"`
	ReviewStatus        int8    `orm:"column(review_status)" description:"订单评价状态"`
	FeedbackStatus      int8    `orm:"column(feedback_status)" description:"订单维权状态"`
	IsEvaluate          int16   `orm:"column(is_evaluate)" description:"是否评价 0为未评价 1为已评价 2为已追评"`
	TaxMoney            float64 `orm:"column(tax_money);digits(10);decimals(2)"`
	ShippingCompanyId   int     `orm:"column(shipping_company_id)" description:"配送物流公司ID"`
	GivePointType       int     `orm:"column(give_point_type)" description:"积分返还类型 1 订单完成  2 订单收货 3  支付订单"`
	PayTime             int     `orm:"column(pay_time);null" description:"订单付款时间"`
	ShippingTime        int     `orm:"column(shipping_time);null" description:"买家要求配送时间"`
	SignTime            int     `orm:"column(sign_time);null" description:"买家签收时间"`
	ConsignTime         int     `orm:"column(consign_time);null" description:"卖家发货时间"`
	CreateTime          int     `orm:"column(create_time);null" description:"订单创建时间"`
	FinishTime          int     `orm:"column(finish_time);null" description:"订单完成时间"`
	IsDeleted           int     `orm:"column(is_deleted)" description:"订单是否已删除"`
	OperatorType        int     `orm:"column(operator_type)" description:"操作人类型  1店铺  2用户"`
	OperatorId          int     `orm:"column(operator_id)" description:"操作人id"`
	RefundBalanceMoney  float64 `orm:"column(refund_balance_money);digits(10);decimals(2)" description:"订单退款余额"`
	FixedTelephone      string  `orm:"column(fixed_telephone);size(50)" description:"固定电话"`
	DistributionTimeOut string  `orm:"column(distribution_time_out);size(50)" description:"配送时间段"`
}


func (t *NsOrder) TableName() string {
	return "ns_order"
}

func init() {
	orm.RegisterModel(new(NsOrder))
}

// AddNsOrder insert a new NsOrder into database and returns
// last inserted Id on success.
func AddNsOrder(m *NsOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetUniOrderNo()  string {
	orderSn := xid.New()
	o := orm.NewOrm()
	if c,_ := o.QueryTable(new(NsOrder)).Filter("OrderNo",orderSn.String()).Count();  c > 0 {
		return GetUniOrderNo()
	}else{
		return orderSn.String()
	}
}


// GetNsOrderById retrieves NsOrder by Id. Returns error if
// Id doesn't exist
func GetNsOrderById(id int) (v *NsOrder, err error) {
	o := orm.NewOrm()
	v = &NsOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNsOrder retrieves all NsOrder matches certain condition. Returns empty list if
// no records exist
func GetAllNsOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsOrder))
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

	var l []NsOrder
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

// UpdateNsOrder updates NsOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsOrderById(m *NsOrder) (err error) {
	o := orm.NewOrm()
	v := NsOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsOrder deletes NsOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsOrder(id int) (err error) {
	o := orm.NewOrm()
	v := NsOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
