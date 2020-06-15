package arrayt

import "testing"

func TestBasicUsage(t *testing.T) {
	// 声明一个map
	m1 := make(map[int]string)
	// 修改值
	m1[1] = "hello"
	m1[2] = "world"
	m1[3] = "!"
	t.Log(m1, len(m1))

	// 获取值
	if v, ok := m1[2]; ok {
		t.Logf("the m1 has the key: 2 and the value is %v", v)
	}

	// 遍历
	for key, value := range m1 {
		t.Logf("key: %v, value: %v", key, value)
	}

	// 删除
	delete(m1, 3)
	t.Log("deleted value:", m1)

	// 初始化声明
	m2 := map[string]string{"name": "Jimmy", "city": "Beijing"}
	t.Log("m2:", m2)
}
