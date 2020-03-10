package models
import (

	"github.com/astaxie/beego/orm"
	"github.com/rs/xid"
)
type NsOrderPayment struct {
	Id int `orm:"-"`
	OutTradeNo    string  `orm:"column(out_trade_no);size(30)" description:"支付单编号"`
	ShopId        int     `orm:"column(shop_id)" description:"执行支付的相关店铺ID（0平台）"`
	Type          int     `orm:"column(type)" description:"订单类型1.商城订单2.交易商支付"`
	TypeAlisId    int     `orm:"column(type_alis_id)" description:"订单类型关联ID"`
	PayBody       string  `orm:"column(pay_body);size(255)" description:"订单支付简介"`
	PayDetail     string  `orm:"column(pay_detail);size(1000)" description:"订单支付详情"`
	PayMoney      float64 `orm:"column(pay_money);digits(10);decimals(2)" description:"支付金额"`
	PayStatus     int8    `orm:"column(pay_status)" description:"支付状态"`
	PayType       int     `orm:"column(pay_type)" description:"支付方式"`
	CreateTime    int     `orm:"column(create_time);null" description:"创建时间"`
	PayTime       int     `orm:"column(pay_time);null" description:"支付时间"`
	TradeNo       string  `orm:"column(trade_no);size(30)" description:"交易号，支付宝退款用，微信传入空"`
	OriginalMoney float64 `orm:"column(original_money);digits(10);decimals(2)" description:"原始支付金额"`
	BalanceMoney  float64 `orm:"column(balance_money);digits(10);decimals(2)" description:"使用余额"`
}

func (t *NsOrderPayment) TableName() string {
	return "ns_order_payment"
}

func init() {
	orm.RegisterModel(new(NsOrderPayment))
}

func GetUniOutTradeNo()  string {
	uuid := xid.New()
	o := orm.NewOrm()
	if c,_ := o.QueryTable(new(NsOrderPayment)).Filter("OutTradeNo",uuid.String()).Count();  c > 0 {
		return GetUniOutTradeNo()
	}else{
		return uuid.String()
	}
}