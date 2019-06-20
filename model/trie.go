package model

import "github.com/jinzhu/gorm"

type Trie struct {
	gorm.Model
	TrieTypeId int64  `gorm:"column:trie_type_id"`
	Key        string `gorm:"column:key"`
	Pv         int64 `gorm:"column:pv"`
	//站在Trie的角度上,外键是TrieType的ID属性，Trie的TrieTypeId作为关联的建
	TrieType TrieType `gorm:"foreignkey:ID;AssociationForeignKey:TrieTypeId"`
}

//表的名字
func (a Trie) TableName() string {
	return "trie"
}
