package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type SysAlbumPicture struct {
	Id            int    `orm:"column(pic_id);auto" description:"相册图片表id"`
	ShopId        uint   `orm:"column(shop_id);null" description:"所属实例id"`
	AlbumId       uint   `orm:"column(album_id)" description:"相册id"`
	IsWide        int    `orm:"column(is_wide)" description:"是否宽屏"`
	PicName       string `orm:"column(pic_name);size(100)" description:"图片名称"`
	PicTag        string `orm:"column(pic_tag);size(255)" description:"图片标签"`
	PicCover      string `orm:"column(pic_cover);size(255)" description:"原图图片路径"`
	PicSize       string `orm:"column(pic_size);size(255)" description:"原图大小"`
	PicSpec       string `orm:"column(pic_spec);size(100)" description:"原图规格"`
	PicCoverBig   string `orm:"column(pic_cover_big);size(255)" description:"大图路径"`
	PicSizeBig    string `orm:"column(pic_size_big);size(255)" description:"大图大小"`
	PicSpecBig    string `orm:"column(pic_spec_big);size(100)" description:"大图规格"`
	PicCoverMid   string `orm:"column(pic_cover_mid);size(255)" description:"中图路径"`
	PicSizeMid    string `orm:"column(pic_size_mid);size(255)" description:"中图大小"`
	PicSpecMid    string `orm:"column(pic_spec_mid);size(100)" description:"中图规格"`
	PicCoverSmall string `orm:"column(pic_cover_small);size(255)" description:"小图路径"`
	PicSizeSmall  string `orm:"column(pic_size_small);size(255)" description:"小图大小"`
	PicSpecSmall  string `orm:"column(pic_spec_small);size(255)" description:"小图规格"`
	PicCoverMicro string `orm:"column(pic_cover_micro);size(255)" description:"微图路径"`
	PicSizeMicro  string `orm:"column(pic_size_micro);size(255)" description:"微图大小"`
	PicSpecMicro  string `orm:"column(pic_spec_micro);size(255)" description:"微图规格"`
	UploadTime    int    `orm:"column(upload_time);null" description:"图片上传时间"`
	UploadType    int    `orm:"column(upload_type);null" description:"图片外链"`
	Domain        string `orm:"column(domain);size(255);null" description:"图片外链"`
	Bucket        string `orm:"column(bucket);size(255);null" description:"存储空间名称"`
}

func (t *SysAlbumPicture) TableName() string {
	return "sys_album_picture"
}

func init() {
	orm.RegisterModel(new(SysAlbumPicture))
}
func NewSysAlbumPicture() *SysAlbumPicture {
	return &SysAlbumPicture{}
}
// AddSysAlbumPicture insert a new SysAlbumPicture into database and returns
// last inserted Id on success.
func AddSysAlbumPicture(m *SysAlbumPicture) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetGoodsPicture(ImgIdArray string) (images []string, err error) {
	o := orm.NewOrm()
	//获取图片
	picIds := strings.Split(ImgIdArray, ",")
	fmt.Println(picIds)
	sysAlbumPicture := NewSysAlbumPicture()
	var picture []*SysAlbumPicture
	num,err := o.QueryTable(sysAlbumPicture).Filter("pic_id__in", picIds).All(&picture)
	if num == 0 {
		return nil,err
	}
	//var images []string
	imgUrl := beego.AppConfig.String("ImgUrl")
	for _,v := range picture{
		//fmt.Sprintf("%s%s", imgUrl, v.PicCoverBig)
		v.PicCoverBig = fmt.Sprintf("%s%s", imgUrl, v.PicCoverBig)
		images = append(images,v.PicCoverBig )
	}

	return images,err
}
// GetSysAlbumPictureById retrieves SysAlbumPicture by Id. Returns error if
// Id doesn't exist
func GetSysAlbumPictureById(id int) (v *SysAlbumPicture, err error) {
	o := orm.NewOrm()
	v = &SysAlbumPicture{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSysAlbumPicture retrieves all SysAlbumPicture matches certain condition. Returns empty list if
// no records exist
func GetAllSysAlbumPicture(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysAlbumPicture))
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

	var l []SysAlbumPicture
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

// UpdateSysAlbumPicture updates SysAlbumPicture by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysAlbumPictureById(m *SysAlbumPicture) (err error) {
	o := orm.NewOrm()
	v := SysAlbumPicture{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysAlbumPicture deletes SysAlbumPicture by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysAlbumPicture(id int) (err error) {
	o := orm.NewOrm()
	v := SysAlbumPicture{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysAlbumPicture{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
