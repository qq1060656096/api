# go-api-result

```go
import "github.com/qq1060656096/go-api"

var profile = struct {
Name string
Age  int
}{Name: "赵1", Age: 18}

result := NewResult("200", "ok")
// 普通结果
result.Simple(profile)

list = append(list, profile)
profile.Name = "赵2"
profile.Age = 19
list = append(list, profile)
profile.Name = "赵3"
profile.Age = 20
list = append(list, profile)
// 分页结果
result.Paging(list, 3, 1, 3)

```