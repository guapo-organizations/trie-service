package lib

import (
	"fmt"
	"github.com/derekparker/trie"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/trie-service/model"
	"sync"
)

//字典树管理

var trie_manage *sync.Map

//单例模式
func init() {
	//初始化
	if trie_manage == nil {
		trie_manage = &sync.Map{}
	}
}

//初始化所有的字典树
func InitTrit() {
	//分页初始化字典树
	db := database.GetMysqlDB()
	var where_id uint
	where_id = 0
	for {
		var trie_model []model.Trie
		//预加载属性
		db.Preload("TrieType").Where("id > ?", where_id).Limit(100).Find(&trie_model)

		if trie_model == nil || len(trie_model) == 0 {
			//没有数据了，分页结束
			return;
		}

		//初始化字典树
		for _, item := range trie_model {

			typeTritInitByIdAndDescribe(TrieType{
				TrieId:   item.TrieTypeId,
				Describe: "",
			}, item.Key)

			//设置id，用于分页
			where_id = item.ID
		}
	}
}

//初始化一颗字典树
func typeTritInitByIdAndDescribe(trie_type TrieType, key string) {

	dict_trie, ok := trie_manage.Load(TrieTypeUniq(trie_type))
	if !ok {
		//没找到这个字典树
		//初始化这棵树
		init_dict_trie := trie.New()
		init_dict_trie.Add(key, "")
		trie_manage.Store(TrieTypeUniq(trie_type), init_dict_trie)
		return
	}
	//字典树存在，添加字符
	dict_trie.(*trie.Trie).Add(key, "")
}

//搜索关键词提示
//trie_id 和 describe 用于标识是那个字典树的
//key 是搜索的关键词
//lexicon_enroll 关键字是否开启录入关键字功能
//callback 是词库录入的回调
func Search(trie_type TrieType, key string, lexicon_enroll bool, callback LexiconEnrollCallBack) ([]string, error) {
	//获取字典树
	dict_trie, ok := trie_manage.Load(TrieTypeUniq(trie_type))
	if !ok {
		return nil, fmt.Errorf("找不到字典树")
	}

	//是否开启词库录取
	if lexicon_enroll {
		//开启了词库录入
		go callback(trie_type, key)
	}
	//查询结果返回
	return dict_trie.(*trie.Trie).FuzzySearch(key), nil
}
