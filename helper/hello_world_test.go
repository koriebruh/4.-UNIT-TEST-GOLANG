package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

/*
--> go test 			  			: run all test and show result test
--> go test -v            			: run all test , show result test and show verbose
--> go test -v - run=TestNameFunc   : select Test func u want
--> go test -v ./...      			: run all package from top root
*/

/* menggagalkan unit test
<-- Fail()      = menandai sebuah unit test gagal namun kode dibawahnya teatap di eksecute dan tidak menghentikan program
<-- FailNow()   = jika menamukan sebuah unit test gagal langusng berhenti
<-- Error()     = seperti logprint error setelah melakukan logprint err akan melakuakn Fail()
<-- Fatal()     = seperti logprint error setelah melakukan logprint err akan melakuakn FailNow()
*/

func TestHelloWorldShank(t *testing.T) { // <-- Fail()
	result := HelloWorld("Shank")
	if result != "Hello Shank" {
		t.Fail()
	}
	fmt.Println("DIKESEKUSI")
}

func TestHelloWorldTeach(t *testing.T) { // <-- FailNow()
	result := HelloWorld("Teach")
	if result != "Hello Teach" {
		t.FailNow()
	}
	fmt.Println("TIDAK DIEKSEKUSI Jika error")
}

func TestHelloWorldBuggy(t *testing.T) { // <-- Error(*arg)
	result := HelloWorld("Buggy")
	if result != "Hello Buggy" {
		t.Error("Error bg harus nya 'Hello Buggy' ")
	}
	fmt.Println("DIEKSEKUSI")
}

func TestHelloWorldKid(t *testing.T) { // <-- Fatal(*arg)
	result := HelloWorld("Kid")
	if result != "Hello Kid" {
		t.Fatal("Error bg harus nya 'Hello Kid' ")
	}
	fmt.Println("TIDAK DIEKSEKUSI Jika error")
}

/* <--- Assert, pengecekan menggantikan if else dan jika penggecekan gagal akan memanggil Fail() sedangkan
Require sama namun jika gagal akan memanggil FailNow()
contoh method yang bisa di panggil assert dan require :

- assert.True(), 	  - assert.False()
- assert.Equal(),     - assert.NotEqual()
- assert.Same()		  - assert.NotSame()
- assert.Nil()		  - assert.NotNil()
- assert.IsType()
- assert.Panics()	  - assert.NotPanics()
- assert.Contains()   - assert.NotContains()
- assert.Len()
- assert.FileExists() - assert.NoFileExists()
- assert.Condition()
*/

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("sipp")
	assert.Equal(t, "Hello sipp", result, "result salah") // <-- expected hasil yang di harapkan, di bandikan dengan result, dan pesan
	fmt.Println("Hahah ini tetap di eksekusi MESKIPUN ERROR")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("leq")
	require.Equal(t, "Hello leq", result, "result salah")
	fmt.Println("GAGAL DI EKSEKUSI JIKA TERDAPAT ERROR")
}

// <-- Skip Test kode dibawahnya tidakakan dijalankan
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" { // <-- darwin itu mama inti sistem operasi MacOs
		t.Skip("Unit test ini tidak bisa jalan di Mac")
	}
	// <-- kode dibawah ini akan di skip jika pernyataan di atas benar
	result := HelloWorld("Iroha")
	require.Equal(t, "Hello Iroha", result, "expected lu salah bg")

}

// <-- before and after test main,
func TestMain(m *testing.M) {
	// <-- before, bisa di manfaatkan misal connection ke data base
	fmt.Println("kode ini akan di eksekusi sebelum unit test di jalankan")

	m.Run() // <-- mengeksekusi semua unit test
	// <-- after
	fmt.Println("Kode ini akan di eksekusi setelah semua unit test selesai")
}

// <-- Sub Test, mendukung pembuatkan func unit test di dalam func unit test mengunakan t.RUN()
func TestSubTest(t *testing.T) { // <-- menjalankan manual semua unit test didalam nya  --> go test -v -run=TestSubTest
	t.Run("Irohana", func(t *testing.T) {
		result := HelloWorld("Irohana")
		require.Equal(t, "Hello Irohana", result, "kok bisa salah?")
	})

	t.Run("Jamal", func(t *testing.T) { // <-- menjalankan unit test ini saja  -->  go test -v -run=TestSubTest/Jamal
		result := HelloWorld("Jamal")
		require.Equal(t, "Hello Jamal", result, "kok bisa salah?")
	})
}

/*
<-- table Test, untuk mengatasi kode yang sama namun value parameter yang berbeda
caranya membuat slice struct yang berisi name,request dam expected lalu gunakan sub test untuk iterasi kode nya
*/
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "HelloWold(Rae)", // <-- ini nama func
			Request:  "Rae",            // <-- ini param
			Expected: "Hello Rae",      // <-- ini hasil yang diharapkan
		},
		{
			Name:     "HelloWorld(Rei)",
			Request:  "Rei",
			Expected: "Hello Rei",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWorld(test.Request)
			require.Equal(t, test.Expected, result, "u got Error")
		})
	}
}

/*	<-- Mock
teknik membuat obj dari suatu objek yang sulit di testing misal memanggil APIcall ThirdParty service contoh lain,
kita ingin menguji mekanisme Query ke database namun kita tidak mau me runing data base

contoh kasus :
- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

<-- kode terdapat pada folder
entity
	|- category.go
repository
	|- category_repository.go
	|- category_repository_mock.go
service
	|- category_service.go
	|- category_service_test.go
*/

// <-- Benchmark digunakan untuk menguji performa (speed) kode kita dilakukan loop dengan banyak otomatis oleh golang
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Irohana")
	}
}

/*
untuk menjalankan :
--> go test -v -bench=.         						       :> mejalankan semua unit test dan semua benchmark
--> go test -v -run=NotMathUnitTest -bench=.  			       :> menjalankan semua benchmark only
--> go test -v -run=NotMathUnitTest -bench=BenchmarkHelloWorld :> menjalankan salah satu benchmark
--> go test -v -bench=. ./...							       :> menjalankan semua unit dan bench dari top root
--> go test -v -run=TidakAda -bench=. ./...					   :> menjalankan semua bench dari root atas
*/

// <-- subBenchmark
func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Irohana", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Irohana")
		}
	})

	b.Run("Jamal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Jamal")
		}
	})
}

// <-- bancmarkTabel, untuk menguji banckmark terhadap function sama dengan parameter yang berbeda
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "HelloWold(Rae)", // <-- ini nama func
			Request: "Rae",            // <-- ini param
		},
		{
			Name:    "HelloWorld(Irohana)",
			Request: "Irohana",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.Request)
			}
		})
	}
}
