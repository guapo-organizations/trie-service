package lib

import (
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/trie-service/model"
	"github.com/jinzhu/gorm"
	"log"
)

//词库管理

//词库录入回调
type LexiconEnrollCallBack func(trie_type TrieType, key string)

func DefaultLexiconEnroll(trie_type TrieType, key string) {
	db := database.GetMysqlDB()
	var trie_type_model *model.TrieType
	db.Where("id = ?", trie_type.TrieId).First(trie_type_model)
	//词库是否注册
	if trie_type_model == nil {
		log.Printf("词库类型没有注册,词库id:%d,词库描述:%s", trie_type.TrieId, trie_type.Describe)
		return
	}

	//关键词是否存在
	var trie_model *model.Trie
	db.Where("trie_type_id = ? AND key = ?", trie_type_model.ID, key).First(trie_model)
	if trie_model == nil {
		//第一次出现该词语，录入
		trie_model := new(model.Trie)
		trie_model.Key = key
		trie_model.TrieTypeId = trie_type.TrieId
		db.Create(trie_model)
		return
	}
	//关键词不是第一次出现，关键词pv + 1
	db.Model(trie_model).Update("pv", gorm.Expr("pv + ?", 1))
}
