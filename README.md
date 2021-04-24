# json进行Unmarshal方法时，根据default的Tag设置默认值

### Example

```go

func main() {
	type user struct {
		Name  string `json:"name" default:"zhouyu"`
		Phone string `json:"phone" default:"123456"`
		Age   int    `json:"age" default:"28"`
	}
	data := []byte(`{"phone":"1572681"}`)
	u := user{}
	err := UnmarshalDefault(data, &u)
	fmt.Println(err, u)
}

```