package controllers

import (
	"encoding/json"
	"errors"
	"github.com/OnfireMrHuang/myblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

func (c *CategoryController) Create() {
	result := models.NewCommResult()
	var v models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddCategory(&v); err == nil {
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

func (c *CategoryController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetCategoryById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *CategoryController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}

	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}

	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error:invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCategory(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func (c *CategoryController) Update() {
	result := models.NewCommResult()
	v := models.Category{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateCategoryById(&v); err == nil {
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

func (c *CategoryController) Delete() {

	result := models.NewCommResult()
	v := models.Category{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.DeleteCategory(v.Id); err == nil {
			c.Ctx.Output.SetStatus(201)
			result.Msg = "OK"
		} else {
			result.Msg = err.Error()
		}
	}
	c.Data["json"] = result
	c.ServeJSON()
}
