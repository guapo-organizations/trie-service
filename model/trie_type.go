package model

import "github.com/jinzhu/gorm"

type TrieType struct {
	gorm.Model
	TrieDescribe int64  `gorm:"column:trie_describe"`
}

//表的名字
func (a TrieType) TableName() string {
	return "trie_type"
}
