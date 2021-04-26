package json

import (
	"fmt"
	"testing"
)

func Test_UnmarshalDefaultWithNoInput(t *testing.T) {
	type user struct {
		Sli   []string `json:"sli"`
		Name  string   `json:"name" default:"-1" `
		Name2 string   `json:"name2" default:"-1"`
		Name3 string   `json:"name3" default:"-1"`
		Name4 string   `json:"name4"`
		Age   int      `json:"age" default:"-1"`
		Age2  int64    `json:"age2" default:"-1"`
		Age3  int64    `json:"age3" default:"-1"`
		Age4  int64    `json:"age4"`
	}
	data := []byte(`{"sli":["1","2"],"name2":"1572681","name3":"","age2":28,"age3":0}`)
	u := new(user)
	err := DefaultWithNoInput(data, u)
	fmt.Println(err, u)
	if err != nil && len(u.Sli) != 0 || u.Name != "-1" || u.Name2 != "1572681" || u.Name3 != "" || u.Name4 != "" || u.Age != -1 || u.Age2 != 28 || u.Age3 != 0 || u.Age4 != 0 {
		t.Fail()
	}
}


func Test_UnmarshalDefaultWithEmpty(t *testing.T) {
	type user struct {
		Sli   []string `json:"sli"`
		Name  string   `json:"name" default:"-1" `
		Name2 string   `json:"name2" default:"-1"`
		Name3 string   `json:"name3" default:"-1"`
		Name4 string   `json:"name4"`
		Age   int      `json:"age" default:"-1"`
		Age2  int64    `json:"age2" default:"-1"`
		Age3  int64    `json:"age3" default:"-1"`
		Age4  int64    `json:"age4"`
	}
	data := []byte(`{"sli":["1","2"],"name2":"1572681","name3":"","age2":28,"age3":0}`)
	u := new(user)
	err := DefaultWithEmpty(data, u)
	fmt.Println(err, u)
	if err != nil && len(u.Sli) != 0 || u.Name != "-1" || u.Name2 != "1572681" || u.Name3 != "-1" || u.Name4 != "" || u.Age != -1 || u.Age2 != 28 || u.Age3 != -1 || u.Age4 != 0 {
		t.Fail()
	}
}