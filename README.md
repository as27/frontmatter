# frontmatter
Parses an io.Reader and extracts the frontmatter 

## Syntax

When you need to store some additional information inside a text you can use a frontmatter logic for doing this. That section ist defines by a delimiter, which should be not used inside the text. Inside that section you can define simple variables with values. 

    +++
    key1 = value1
    key 2 = another seconde value
    +++
    # Some header here

    some more text here

So the delimiter here is `+++` that part will be extracted inside a map.

```go
doc := `+++
key1 = first key
key2 = 2ndKey 
+++
# My Document
...
`
b := &bytes.Buffer{}
fm, _ := Parse("+++", strings.NewReader(doc), b)
fmt.Printf("%#v\n", fm)
fmt.Println(b)
// Output: map[string]string{"key1":"first key", "key2":"2ndKey"}
// # My Document
// ...
```