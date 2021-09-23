package cogen_test

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

func ExampleMySQL() {
	g, err := cogen.MySQL(Item{})
	if err != nil {
		panic(err)
	}
	fmt.Println(g.String())
	// Output: CREATE TABLE `t_item` (`id` bigint(20) unsigned NOT NULL,`name` varchar(64) NOT NULL DEFAULT '',`price` int unsigned NOT NULL DEFAULT 0,PRIMARY KEY (`id`),UNIQUE INDEX idx_t_item_name (`name`))
}
