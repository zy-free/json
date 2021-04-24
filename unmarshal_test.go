package json

import "testing"

func Test_Unmarshal(t *testing.T) {
	type user struct {
		Name  string `json:"name" default:"zhouyu"`
		Phone string `json:"phone" default:"123456"`
		Age   int    `json:"age" default:"28"`
	}
	data := []byte(`{"phone":"1572681"}`)
	u := user{}
	err := UnmarshalDefault(data, &u)

	if err != nil && u.Phone != "1572681" || u.Name != "zhouyu" || u.Age != 28 {
		t.Fail()
	}
}
