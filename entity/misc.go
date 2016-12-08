package entity

type RedisTest struct {
	Int    int
	String string
	Bool   bool
	Array  []int
	Struct
}

type Struct struct {
	Name string
	Age  int
}
