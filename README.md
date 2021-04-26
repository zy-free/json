# json进行Unmarshal方法时，根据default的Tag设置默认值
### DefaultWithNoInput方法
    当没有输入时设置默认值，适合httpServer的处理，判断前端是否输入来做增量更新
### DefaultWithEmpty方法
    当jsonValue为零值时，设置默认值
### Example

```go

import (
	"fmt"
	xjson "github.com/zy-free/json"
)

func main() {
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
	u := user{}
	err := xjson.DefaultWithNoInput(data, &u)
	fmt.Println(err, u)
    // u.Name != "-1" 
    // u.Name2 != "1572681" 
    // u.Name3 != "" 
    // u.Name4 != "" 
    // u.Age != -1 
    // u.Age2 != 28 
    // u.Age3 != 0 
    // u.Age4 != 0
    err = DefaultWithEmpty(data, &u)
	fmt.Println(err, u)
    // u.Name != "-1" 
    // u.Name2 != "1572681" 
    // u.Name3 != "-1" 
    // u.Name4 != ""
    // u.Age != -1
    // u.Age2 != 28
    // u.Age3 != -1
    // u.Age4 != 0
}


```