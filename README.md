# cogen

> *co-generator*, or *codegen*.

*cogen* is a DDL codegen for GORM structs.

## Example

```go
package main

import (
	"fmt"

	"github.com/anqur/cogen"
)

type Item struct {
	ID    uint64 `gorm:"column:id;primary_key;auto_increment;not null;type:bigint(20) unsigned"`
	Name  string `gorm:"column:name;not null;type:varchar(64);default:\"\";uniqueindex"`
	Price uint32 `gorm:"column:price;not null;type:uint;default:0"`
}

func (Item) TableName() string {
	return "t_item"
}

func main() {
	g, err := cogen.MySQL(Item{})
	if err != nil {
		panic(err)
	}
	fmt.Println(g.String())
}
```

## Why

* It's hard to do that in GORM. We need some hacks and tricks
* I believe that DB schemas are mostly not the real business entities without invariants. Let's generate the DDLs from
  some intermediate ORM structs!

## License

MIT
