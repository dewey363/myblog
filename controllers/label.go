package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/sinksmell/LanBlog/models"
	"strconv"
	"strings"
)

type labelResult struct {
	Success bool `json:"success"`
}

type LabelController struct {
	beego.Controller
}

func (c *LabelController) URLMapping() {
	c.Mapping("Create",c.Create)
	c.Mapping("GetOne",c.GetOne)
	c.Mapping("GetAll",c.GetAll)
	c.Mapping("Update",c.Update)
	c.Mapping("Delete",c.Delete)
}

func (c *LabelController) Create() {
	var v models.Label
	result := models.NewCommResult()
	if err :=json.Unmarshal(c.Ctx.Input.RequestBody,&v);err == nil {
		if _,err := models.AddLabel(&v);err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	} else {
		result.Msg = err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LabelController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idStr)
	v,err := models.GetLabelById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *LabelController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v:= c.GetString("fields");v != "" {
		fields= strings.Split(v,",")
	}

	if v,err:= c.GetInt64("limit");err == nil {
		limit = v
	}

	if v,err := c.GetInt64("offset");err == nil {
		offset = v
	}

	if v := c.GetString("sortby");v != "" {
		sortby = strings.Split(v,",")
	}

	if v := c.GetString("order"); v!= "" {
		order = strings.Split(v,",")
	}

	if v:= c.GetString("query");v != "" {
		for _,cond := range strings.Split(v,",") {
			kv := strings.SplitN(cond,":",2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k,v := kv[0],kv[1]
			query[k] =v
		}
	}

	l,err := models.GetAllLabel(query,fields,sortby,order,offset,limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (c *LabelController) Update()  {
	var v models.Label
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&v);err == nil {
		if err := models.UpdateLabelById(&v);err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	} else {
		result.Msg = err.Error()
	}
	c.Data["json"] = result
	c.ServeJSON()
}

func (c *LabelController) Delete() {
	var v models.Label
	result := models.NewCommResult()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&v); err == nil {

		if err := models.DeleteLabel(v.Id); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}