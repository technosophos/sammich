//go:generate ../sammich smap int *MyType
package main

type MyType struct {
	Name string
}

func main() {
	mt := &MyType{"Hello"}
	mm := NewIntMyTypeMap()

	mm.Put(1, mt)
	mm.Put(2, mt)

	t1, ok := mm.Get(1)
	if !ok {
		panic("Could not find one")
	}
	t2, _ := mm.Get(2)

	if t1.Name != t2.Name {
		panic("expected same mt")
	}

	mm.Delete(2)

	println(t1.Name)
}
