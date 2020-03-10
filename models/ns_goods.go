package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type NsGoods struct {
	Id                int     `orm:"column(goods_id);auto" description:"商品id(SKU)"`
	GoodsName         string  `orm:"column(goods_name);size(100)" description:"商品名称"`
	ShopId            uint    `orm:"column(shop_id)" description:"店铺id"`
	CategoryId        uint    `orm:"column(category_id)" description:"商品分类id" json:"-"`
	CategoryId1       uint    `orm:"column(category_id_1)" description:"一级分类id"`
	CategoryId2       uint    `orm:"column(category_id_2)" description:"二级分类id"`
	CategoryId3       uint    `orm:"column(category_id_3)" description:"三级分类id"`
	BrandId           uint    `orm:"column(brand_id)" description:"品牌id"`
	GroupIdArray      string  `orm:"column(group_id_array);size(255)" description:"店铺分类id 首尾用,隔开"`
	PromotionType     int8    `orm:"column(promotion_type)" description:"促销类型 0无促销，1团购，2限时折扣"`
	PromoteId         int     `orm:"column(promote_id)" description:"促销活动ID"`
	GoodsType         int8    `orm:"column(goods_type)" description:"实物或虚拟商品标志 1实物商品 0 虚拟商品 2 F码商品"`
	MarketPrice       float64 `orm:"column(market_price);digits(10);decimals(2)" description:"市场价"`
	Price             float64 `orm:"column(price);digits(19);decimals(2)" description:"商品原价格"`
	PromotionPrice    float64 `orm:"column(promotion_price);digits(10);decimals(2)" description:"商品促销价格"`
	PointExchangeType int8    `orm:"column(point_exchange_type)" description:"积分兑换类型 0 非积分兑换 1 只能积分兑换 "`
	PointExchange     int     `orm:"column(point_exchange)" description:"积分兑换"`
	GivePoint         int     `orm:"column(give_point)" description:"购买商品赠送积分"`
	IsMemberDiscount  int     `orm:"column(is_member_discount)" description:"参与会员折扣"`
	ShippingFee       float64 `orm:"column(shipping_fee);digits(10);decimals(2)" description:"运费 0为免运费"`
	ShippingFeeId     int     `orm:"column(shipping_fee_id)" description:"售卖区域id 物流模板id  ns_order_shipping_fee 表id"`
	Stock             int     `orm:"column(stock)" description:"商品库存"`
	MaxBuy            int     `orm:"column(max_buy)" description:"限购 0 不限购"`
	Clicks            uint    `orm:"column(clicks)" description:"商品点击数量"`
	//MinStockAlarm       int     `orm:"column(min_stock_alarm)" description:"库存预警值"`
	Sales        uint   `orm:"column(sales)" description:"销售数量"`
	Collects     uint   `orm:"column(collects)" description:"收藏数量"`
	Star         uint8  `orm:"column(star)" description:"好评星级"`
	Evaluates    uint   `orm:"column(evaluates)" description:"评价数"`
	Shares       int    `orm:"column(shares)" description:"分享数"`
	ProvinceId   uint   `orm:"column(province_id)" description:"一级地区id"`
	CityId       uint   `orm:"column(city_id)" description:"二级地区id"`
	Picture      int    `orm:"column(picture)" description:"商品主图"`
	Keywords     string `orm:"column(keywords);size(255)" description:"商品关键词"`
	Introduction string `orm:"column(introduction);size(255)" description:"商品简介，促销语"`
	Description  string `orm:"column(description)" description:"商品详情"`
	QRcode       string `orm:"column(QRcode);size(255)" description:"商品二维码"`
	//Code                string  `orm:"column(code);size(50)" description:"商家编号"`
	IsStockVisible   int     `orm:"column(is_stock_visible)" description:"页面不显示库存"`
	IsHot            int     `orm:"column(is_hot)" description:"是否热销商品"`
	IsRecommend      int     `orm:"column(is_recommend)" description:"是否推荐"`
	IsNew            int     `orm:"column(is_new)" description:"是否新品"`
	IsPreSale        int     `orm:"column(is_pre_sale);null"`
	IsBill           int     `orm:"column(is_bill)" description:"是否开具增值税发票 1是，0否"`
	State            int8    `orm:"column(state)" description:"商品状态 0下架，1正常，10违规（禁售）"`
	Sort             int     `orm:"column(sort)" description:"排序"`
	ImgIdArray       string  `orm:"column(img_id_array);size(1000);null" description:"商品图片序列"`
	SkuImgArray      string  `orm:"column(sku_img_array);size(1000);null" description:"商品sku应用图片列表  属性,属性值，图片ID"`
	MatchPoint       float32 `orm:"column(match_point);null" description:"实物与描述相符（根据评价计算）"`
	MatchRatio       float32 `orm:"column(match_ratio);null" description:"实物与描述相符（根据评价计算）百分比"`
	GoodsAttributeId int     `orm:"column(goods_attribute_id)" description:"商品类型"`
	GoodsSpecFormat  string  `orm:"column(goods_spec_format)" description:"商品规格"`
	GoodsWeight      float64 `orm:"column(goods_weight);digits(8);decimals(2)" description:"商品重量"`
	GoodsVolume      float64 `orm:"column(goods_volume);digits(8);decimals(2)" description:"商品体积"`
	//ShippingFeeType     int     `orm:"column(shipping_fee_type)" description:"计价方式1.重量2.体积3.计件"`
	//SupplierId          int     `orm:"column(supplier_id)" description:"供货商id"`
	//SaleDate            int     `orm:"column(sale_date);null" description:"上下架时间"`
	//CreateTime          int     `orm:"column(create_time);null" description:"商品添加时间"`
	//UpdateTime          int     `orm:"column(update_time);null" description:"商品编辑时间"`
	//MinBuy              int     `orm:"column(min_buy)" description:"最少买几件"`
	//VirtualGoodsTypeId  int     `orm:"column(virtual_goods_type_id);null" description:"虚拟商品类型id"`
	//ProductionDate      int     `orm:"column(production_date)" description:"生产日期"`
	ShelfLife         string `orm:"column(shelf_life);size(50)" description:"保质期"`
	GoodsVideoAddress string `orm:"column(goods_video_address);size(455);null" description:"商品视频地址，不为空时前台显示视频"`
	//PcCustomTemplate    string  `orm:"column(pc_custom_template);size(255)" description:"pc端商品自定义模板"`
	//WapCustomTemplate   string  `orm:"column(wap_custom_template);size(255)" description:"wap端商品自定义模板"`
	//MaxUsePoint         int     `orm:"column(max_use_point)" description:"积分抵现最大可用积分数 0为不可使用"`
	//IsOpenPresell       int8    `orm:"column(is_open_presell)" description:"是否支持预售"`
	//PresellTime         int     `orm:"column(presell_time)" description:"预售发货时间"`
	//PresellDay          int     `orm:"column(presell_day)" description:"预售发货天数"`
	//PresellDeliveryType int     `orm:"column(presell_delivery_type)" description:"预售发货方式1. 按照预售发货时间 2.按照预售发货天数"`
	//PresellPrice        float64 `orm:"column(presell_price);digits(10);decimals(2)" description:"预售金额"`
	//GoodsUnit           string  `orm:"column(goods_unit);size(20)" description:"商品单位"`
	//IntegralGiveType    int     `orm:"column(integral_give_type)" description:"积分赠送类型 0固定值 1按比率"`
}
type NsGoodsResult struct {
	Id        int    `orm:"column(goods_id);auto" description:"商品id(SKU)"`
	GoodsName string `orm:"column(goods_name);size(100)" description:"商品名称"`
}

func (t *NsGoods) TableName() string {
	return "ns_goods"
}

func init() {
	orm.RegisterModel(new(NsGoods))
	orm.RegisterModel(new(NsGoodsResult))
}
func GetHomeGoodsList(offset int64, limit int64) (list *[]orm.Params, totalCount int64, err error) {
	o := orm.NewOrm()
	//var goods []*Goods
	//num, err := o.QueryTable("goods").Filter("Tags__Tag__Name", "golang").All(&goods)
	sql := "select a.goods_id,a.goods_name,a.price,a.market_price,a.sales,b.pic_cover_small from ns_goods a  left join sys_album_picture b on a.picture=b.pic_id " +
		"where a.is_hot=1 and a.stock>0 order by a.sales desc limit ?,? "
	var maps []orm.Params
	_, err = o.Raw(sql, offset, limit).Values(&maps)
	imgUrl := beego.AppConfig.String("ImgUrl")
	for _, v := range maps {
		v["pic_cover_small"] = fmt.Sprintf("%s%s", imgUrl, v["pic_cover_small"])
	}
	if err != nil {
		return nil, 0, errors.New("暂无数据")

	}
	sqlCount := "select count(*)  from ns_goods where is_hot=1 and stock>0 "
	err = o.Raw(sqlCount).QueryRow(&totalCount)
	if err != nil {
		return nil, 0, errors.New("暂无数据")

	}

	return &maps, totalCount, nil
}

// AddNsGoods insert a new NsGoods into database and returns
// last inserted Id on success.
func AddNsGoods(m *NsGoods) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNsGoodsById retrieves NsGoods by Id. Returns error if
// Id doesn't exist
func GetNsGoodsById(id int) (goods *NsGoods, err error) {
	o := orm.NewOrm()
	g := NsGoods{Id: id}
	//var goods orm.Params
	if err = o.Read(&g); err == nil {
		imgUrl := beego.AppConfig.String("ImgUrl")
		g.QRcode = fmt.Sprintf("%s%s", imgUrl, g.QRcode)

		return &g, nil
	}
	return nil, err
}



// GetAllNsGoods retrieves all NsGoods matches certain condition. Returns empty list if
// no records exist
func GetAllNsGoods(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NsGoods))
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

	var l []NsGoods
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

// UpdateNsGoods updates NsGoods by Id and returns error if
// the record to be updated doesn't exist
func UpdateNsGoodsById(m *NsGoods) (err error) {
	o := orm.NewOrm()
	v := NsGoods{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNsGoods deletes NsGoods by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNsGoods(id int) (err error) {
	o := orm.NewOrm()
	v := NsGoods{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NsGoods{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
