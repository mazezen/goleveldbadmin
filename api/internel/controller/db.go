package controller

import (
	"encoding/json"
	"github.com/mazezen/goleveldbadmin/framework"
	"github.com/mazezen/goleveldbadmin/service"
	"io"
	"net/http"
)

type DbController struct{}

var dbService = new(service.DbService)

func (d *DbController) Router(router *framework.RouterHandler) {
	router.Router("/lb/findAll", d.findAll)
	router.Router("/lb/addData", d.add)
	router.Router("/lb/dataInfo", d.info)
	router.Router("/lb/deleteData", d.delete)
	router.Router("/lb/clear", d.clear)
}

type findAllPayload struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Key      string `json:"key"`
}

func (d *DbController) findAll(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var result = &framework.Result{
		Code:    2000,
		Message: "成功",
	}

	var payload findAllPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}
	count, res := dbService.FindAll(payload.Page, payload.PageSize, payload.Key)
	var m = make(map[string]interface{})
	m["total"] = count
	m["list"] = res

	result.Data = m
	framework.JsonOk(w, result)
}

type addPayload struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (d *DbController) add(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var result = &framework.Result{
		Code:    2000,
		Message: "成功",
	}
	var payload addPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		if err == io.EOF {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		} else {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		}
	}
	if err := dbService.Add(payload.Key, payload.Value); err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}
	framework.JsonOk(w, result)
}

func (d *DbController) info(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var result = &framework.Result{
		Code:    2000,
		Message: "成功",
	}

	var payload addPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		if err == io.EOF {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		} else {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		}
	}

	res, err := dbService.Info(payload.Key)
	if err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}
	result.Data = res
	framework.JsonOk(w, result)
}

func (d *DbController) delete(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var result = &framework.Result{
		Code:    2000,
		Message: "成功",
	}

	var payload addPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		if err == io.EOF {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		} else {
			result.Code = 5000
			result.Message = err.Error()
			framework.JsonOk(w, result)
			return
		}
	}

	err = dbService.Delete(payload.Key)
	if err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}
	framework.JsonOk(w, result)
}

func (d *DbController) clear(w http.ResponseWriter, r *http.Request) {
	cors(w, r)
	var result = &framework.Result{
		Code:    2000,
		Message: "成功",
	}
	if err := dbService.Clear(); err != nil {
		result.Code = 5000
		result.Message = err.Error()
		framework.JsonOk(w, result)
		return
	}
	framework.JsonOk(w, result)
}
