# json进行Unmarshal方法时，根据default的Tag设置默认值

### Example

```go

import (
	"fmt"
	xjson "github.com/zy-free/json"
)

func main() {
	type user struct {
		Name  string `json:"name" default:"zhouyu"`
		Phone string `json:"phone" default:"123456"`
		Age   int    `json:"age" default:"28"`
	}
	data := []byte(`{"phone":"1572681"}`)
	u := user{}
	err := xjson.UnmarshalDefault(data, &u)
	fmt.Println(err, u)
}


```