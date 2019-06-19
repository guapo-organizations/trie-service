package model

import "github.com/jinzhu/gorm"

type Trie struct {
	gorm.Model
	TrieTypeId int64  `gorm:"column:trie_type_id"`
	Key        string `gorm:"column:key"`
	Pv         string `gorm:"column:pv"`
}

//表的名字
func (a Trie) TableName() string {
	return "trie"
}
