package database

import (
	"trunk_exem/employment_examination/golang/models"
)

var DB []models.User

func ClosureKVS() func(c string, k string, v interface{}) interface{} {
	init := make(map[string]interface{})
	return func(cmd string, k string, v interface{}) interface{} {
		switch cmd {
		case "ADD":
			init[k] = v
			return nil
		case "GET":
			res := init[k]
			return res
		case "RESET":
			init = make(map[string]interface{})
			return nil
		case "KEYS":
			resp := make([]string, 0)
			for k := range init {
				resp = append(resp, k)
			}
			return resp
		case "VALUES":
			resp := make([]interface{}, 0)
			for _, v := range init {
				resp = append(resp, v)
			}
			return resp
		}
		return nil
	}
}
