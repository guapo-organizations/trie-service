package service

import (
	"context"
	"github.com/guapo-organizations/trie-service/lib"
	"github.com/guapo-organizations/trie-service/proto/trie"
)

//关键词搜索
func (this *TrieService) KeySearch(ctx context.Context, request *trie.KeySearchTrieRequest) (*trie.KeySearchTrieResponse, error) {
	//词库录取回调函数
	var callback lib.LexiconEnrollCallBack
	//判断是否开启词库录取
	if request.KeyEnroll {
		//开启关键字录取
		callback = lib.DefaultLexiconEnroll
	} else {
		//不开启关键词录取
		callback = nil
	}

	//搜索结果
	search_result, err := lib.Search(lib.TrieType{
		TrieId:   request.TrieType.Id,
		Describe: request.TrieType.TrieDescribe,
	}, request.Key, request.KeyEnroll, callback)

	if err != nil {
		return nil, err
	}

	//返回结果
	return &trie.KeySearchTrieResponse{
		KeyList: search_result,
	}, nil
}
