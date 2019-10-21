package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
)

type Category struct {
	Id          int      `json:"id" orm:"column(id);auto"`
	Name        string   `json:"name" orm:"column(name);size(100)"`
	Description string   `json:"description" orm:"column(description);size(500)"`
	Topics      []*Topic `json:"-" orm:"reverse(many)"`
}

func (t *Category) TableName() string {
	return "category"
}

//添加分类
func AddCategory(m *Category) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

//通过ID得到一个分类
func GetCategoryById(id int) (v *Category, err error) {
	o := orm.NewOrm()
	v = &Category{Id: id}
	if err = o.Read(v); err != nil {
		return nil, err
	}
	return v, nil
}

func GetAllCategory(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	//返回QuerySeter对象来组织查询，参数可以是字符串或者结构体
	qs := o.QueryTable(new(Category))
	//通过条件过滤结果
	/*	示例：
		qs.Filter("profile__isnull", true).Filter("name", "slene")
		// WHERE profile_id IS NULL AND name = 'slene'
	*/
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	//根据order by条件，sort是需要排序的字段，order是排序的方式
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order.Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order.Must be either [asc|desc]")
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

	var l []Category
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

func UpdateCategoryById(m *Category) (err error) {
	o := orm.NewOrm()
	v := Category{Id: m.Id}

	if err = o.Read(&v); err != nil {
		return
	}
	var num int64
	if num, err = o.Update(m); err != nil {
		return
	}
	fmt.Println("Number of records updated in database:", num)
	return
}

func DeleteCategory(id int) (err error) {
	o := orm.NewOrm()
	v := Category{Id: id}
	if err = o.Read(&v); err != nil {
		return
	}
	var num int64
	if num, err = o.Delete(&Category{Id: id}); err != nil {
		return
	}
	fmt.Println("Number of records deleted in database:", num)
	return
}
