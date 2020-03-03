package api

import (
	"encoding/json"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"testing"
)

var result = NewResult("200", "ok")
var profile = struct {
Name string
Age  int
}{Name: "赵1", Age: 18}
var profile1 = profile
var list = make([]interface{}, 0)

func init()  {
	list = append(list, profile)
	profile.Name = "赵2"
	profile.Age = 19
	list = append(list, profile)
	profile.Name = "赵3"
	profile.Age = 20
	list = append(list, profile)
}


func TestResultAndToJson(t *testing.T) {
	result := NewResult("200", "ok")
	result.Simple(profile1)
	jsonStr, err := json.Marshal(result)
	assert.Equal(t, nil, err, err)
	expected := `{"code":"200","message":"ok","data":{"Name":"赵1","Age":18}}`
	assert.Equal(t, expected, string(jsonStr))

	result.Paging(list, 3, 1, 3)
	jsonStr, err = json.Marshal(result)
	assert.Equal(t, nil, err, err)
	expected = `{"code":"200","message":"ok","data":{"totalCount":3,"totalPageCount":1,"page":1,"pageSize":3,"lists":[{"Name":"赵1","Age":18},{"Name":"赵2","Age":19},{"Name":"赵3","Age":20}]}}`
	assert.Equal(t, expected, string(jsonStr))
}

func TestResultAndToXml(t *testing.T) {

	result.Simple(profile1)
	xmlStr, err := xml.Marshal(result)
	assert.Equal(t, nil, err, err)
	expected := `<Result><code>200</code><message>ok</message><data><Name>赵1</Name><Age>18</Age></data></Result>`
	assert.Equal(t, expected, string(xmlStr))


	result.Paging(list, 3, 1, 2)
	xmlStr, err = xml.Marshal(result)
	assert.Equal(t, nil, err, err)
	expected = `<Result><code>200</code><message>ok</message><data><totalCount>3</totalCount><totalPageCount>1</totalPageCount><page>1</page><pageSize>2</pageSize><lists><Name>赵1</Name><Age>18</Age></lists><lists><Name>赵2</Name><Age>19</Age></lists><lists><Name>赵3</Name><Age>20</Age></lists></data></Result>`
	assert.Equal(t, expected, string(xmlStr))
}
