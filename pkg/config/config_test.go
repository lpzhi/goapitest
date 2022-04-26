package config

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name":  "lp",
			"debug": true,
		}
	})

	for k, fn := range ConfigFuncs {
		fmt.Println(k)
		mp := fn()
		fmt.Println(mp["name"])
	}

	var i interface{} = []int{1, 23, 9}

	v := reflect.ValueOf(i)

	fmt.Println(reflect.Zero(v.Type()).Interface())
}

type Order struct {
	orderId    int    `json:"order_id" validate:"required"`
	customerId string `json:"customer_id validate:"required"`
}
type Kd uint8

func getKd() Kd {
	return Kd(9)
}

func reflectInfo(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	fmt.Println("Type ", t)
	fmt.Println("Value ", v)

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)
		tag := t.Field(i).Tag.Get("json")
		validate := t.Field(i).Tag.Get("validate")

		fmt.Println(fv, ft, tag, validate)
	}
}

func Test_rf(t *testing.T) {
	o := Order{
		orderId:    999,
		customerId: "lp",
	}
	reflectInfo(o)

	kd := getKd()
	fmt.Println(kd)
	var i interface{}
	i = 100
	val, ok := i.(int)

	fmt.Println()
	fmt.Println(val, ok)
	add1(1, 23, 5)
}

func add1(defaultValue ...interface{}) {
	fmt.Println(defaultValue[0])
}
