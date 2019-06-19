package lib

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

//字典树的类型；这棵字典树是谁的
type TrieType struct {
	TrieId   int64
	Describe string
}

//生成字典树的唯一标识
func TrieTypeUniq(trie_type TrieType) string {
	type_key := strconv.FormatInt(trie_type.TrieId, 10) + trie_type.Describe
	md5_type_key := md5.Sum([]byte(type_key))
	return fmt.Sprintf("%x", md5_type_key)
}
