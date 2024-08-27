package helper

/* RULES
<-- aturan jika ingin membuat test tamabhakan _test jika pada file ini berarti hello_world_test.go
<-- dan aturan untuk func di awali Test
<-- dan aturan untuk parameter (t *testing.T) dan tidak boleh mengembalikan return value
*/

/*
Assertion digunakan untuk menggantikan if else dalam pengecekan data kita perlu menambahkan
library -> go get github.com/stretchr/testify
*/

func HelloWorld(name string) string {
	return "Hello " + name
}
