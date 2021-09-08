package di

import (
	"github.com/go-resty/resty/v2"
	"github.com/mix-go/xdi"
)

func init() {
	obj := xdi.Object{
		Name: "resty",
		New: func() (i interface{}, e error) {
			return resty.New(), nil
		},
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Resty() (s *resty.Client) {
	if err := xdi.Populate("resty", &s); err != nil {
		panic(err)
	}
	return
}
