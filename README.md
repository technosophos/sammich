# Sammich

[[https://xkcd.com/149/]]
Make type-safe collections for me.

This is more a fun proof of concept than anything else. I didn't really
write it for production use.

## Usage

To generate synchronized maps, add a generate header to your source.

```go
//go:generate sammich smap string MyType
package foo

type MyType int

func main() {
  foo.NewStringMyTypeMap()
  foo.Put("key", 7)

  println(foo.Get("key"))
}
```

Then run `go generate` on that file. It will create a file named
`string_mytype_map.go` with the sync map implementation.

### The go generate Header

The header you use in go generate should look like this:

```
//go generate COLLECTION TYPE [TYPE] [EXTRA_PACKAGES...]
```

- Maps take two types: `go:generate smap KeyType ValueType`
- Other collections take only one.
- `EXTRA_PACKAGES` are additional things to drop in the `import` section

Currently supported COLLECTION types are:

- `smap`: Synchronized map.
