# Notes for Go (golang)

## _Go commands_

---

| Command    | What it does                                               |
| ---------- | :--------------------------------------------------------- |
| go build   | compiles a bunch of go source code files                   |
| go run     | compiles and executes one or two files                     |
| go fmt     | formats all the code in each file in the current directory |
| go install | compiles and installs a package                            |
| go get     | downloads the raw source code of someone else's package    |
| go test    | runs any tests associated with the current project         |

---

```
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
```

### Breaking this function down

---

```
func (d deck) toString() string
```

* In this line I'm declaring a function named toString
* (d deck) is a receiver function used to give me access to the deck value "d"
* The last "string" argument is telling Golang that this function is going to return a string

---

```
return strings.Join([]string(d), ",")
```

* I am returning an in-built method called "Join" by using the strings library provided by Golang
* Inside of the "strings.Join()", the argument required is a slice of the type string
* Slice of string requires the argument that needs to be joined. In this case the value "d" of the type deck
* The last and final argument is a separator that tells the join function how to split between indexes of the slice we are trying to join.

### Example

---

```
{"ONE", "TWO", "THREE"}
```

### Would become

---

```
"ONE,TWO,THREE"
```

## Pointers

---

&variable - give me the memory address of the value this variable is pointing at

\*pointer - give me the value this memory address is pointing at

Turn _address_ into _value_ with \*address

Turn _value_ into _address_ with &value

| Value Types | Reference Types |
| ----------- | :-------------- |
| int         | slices          |
| float       | maps            |
| string      | channels        |
| bool        | pointers        |
| structs     | functions       |

use pointers with _VALUE TYPES_

don't worry about pointers when dealing with _REFERENCE TYPES_

## Maps vs Structs

---

| Maps                                                | Structs                                                       |
| --------------------------------------------------- | :------------------------------------------------------------ |
| All _keys_ must be the same type                    | Values can be of different type                               |
| All _values_ must be the same type                  | Keys don't support indexing                                   |
| Keys are indexed - we can iterate over them         | Value Type!                                                   |
| Use to represent a collection of related properties | You need to know all the different fields at compile time     |
| Reference Type!                                     | Use to represent a "thing" with a lot of different properties |
