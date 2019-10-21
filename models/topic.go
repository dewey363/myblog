package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Topic struct {
	Id              int       `json:"id" orm:"column(id):auot"`
	Uid             int64     `json:"uid" orm:"column(uid)"`
	Title           string    `json:"title" orm:"column(title);size(255)"`
	Content         string    `json:"content" orm:"column(content);type(text)"`
	Summary         string    `json:"summary" orm:"column(summary);size(200)"`
	Attachment      string    `json:"url" orm:"column(attachment);size(255)"`
	Category        *Category `json:"cate" orm:"rel(fk);on_delete(do_nothing)"`
	Labels          []*Label  `json:"labels" orm:"rel(m2m)"`
	Created         time.Time `orm:"auto_now_add;column(created);type(datetime)"`
	Updated         time.Time `json:"-" orm:"auto_now;column(updated);type(datetime)"`
	Deleted         time.Time `json:"-" orm:"auto_now;column(deleted);type(datetime)"`
	Views           int64     `json:"views" orm:"column(views)"`
	Author          string    `json:"author" orm:"column(author);size(255)"`
	ReplyTime       time.Time `json:"-" orm:"column(reply_time);type(datetime);null"`
	ReplyCount      int64     `json:"reply_count" orm:"column(reply_count)"`
	ReplyLastUserId int64     `json:"reply_last_user_id" orm:"column(reply_last_user_id)"`
}

func (t *Topic) TableName() string {
	return "topic"
}

func AddTopic(m *Topic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err != nil {
		return id, err
	}
	m2m := o.QueryM2M(m, "Labels")
	_, err = m2m.Add(m.Labels)
	return id, err
}

func GetTopicByID(id int) (v *Topic, err error) {
	o := orm.NewOrm()
	v = &Topic{Id: id}
	if err = o.Read(v); err == nil {
		o.LoadRelated(v, "Category")
		o.LoadRelated(v, "Labels")
		return v, nil
	}
	return nil, err
}

func GetTopicByCateID(cateId int) ([]*Topic, error) {
	o := orm.NewOrm()
	cate := &Category{Id: cateId}
	var topics []*Topic
	if err := o.Read(cate); err != nil {
		return topics, err
	}
	//通过orm提供的自动载入关系字段，实现分类查询
	_, err := o.LoadRelated(cate, "Topics")
	topics = cate.Topics
	for _, v := range topics {
		o.LoadRelated(v, "Category")
		o.LoadRelated(v, "Labels")
	}
	return topics, err
}

func GetTopicByLabelID(tagId int) ([]*Topic, error) {
	o := orm.NewOrm()
	label := &Label{Id: tagId}
	var topics []*Topic
	if err := o.Read(label); err != nil {
		return topics, err
	}
	//通过orm提供的自动载入关系字段，实现分类查询
	_, err := o.LoadRelated(label, "Topics")
	topics = label.Topics
	return topics, err
}

func GetAllTopics(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) ([]*Topic, error) {
	var err error
	o := orm.NewOrm()
	qs := o.QueryTable(new(Topic))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	//order by
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

	var topics []*Topic
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&topics, fields...); err == nil {

		for _, topic := range topics {
			o.Read(topic)
			o.LoadRelated(topic, "Category")
			_, err := o.LoadRelated(topic, "Labels")
			//测试是否正确地载入关系字段
			//bytes, _ := json.Marshal(topic)
			//	beego.BeeLogger.Info("%+v",topic)
			if err != nil {
				return nil, err
			}
		}
		return topics, nil
	}
	return nil, err
}

func UpdateTopicByID(m *Topic) (err error) {
	o := orm.NewOrm()
	v := Topic{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		tags := m.Labels
		m2m := o.QueryM2M(m, "Labels")
		m2m.Clear()
		if num, err = o.Update(m); err == nil {
			m2m.Add(tags)
			fmt.Println("number of records updated in databses:", num)
		}
	}
	return err
}

func DeleteTopic(id int) (err error) {
	o := orm.NewOrm()
	v := Topic{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		m2m := o.QueryM2M(&v, "Labels")
		m2m.Clear()
		if num, err = o.Delete(&Topic{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func FastFind(vagueName string) ([]*Topic, error) {
	var topics_title []*Topic
	var topics_content []*Topic
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	qs.Filter("title__contains", vagueName).All(&topics_title)
	_, err := qs.Filter("content__contains", vagueName).All(&topics_content)
	topics := append(topics_title, topics_content...)
	return topics, err
}
