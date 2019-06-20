# 字典树服务

字典树，项目的字典树


# 服务架构采用自己封装的


# 开启grpc网关之后

## 参数解释

trie_type.id 需要在数据库配置

trie_type.trie_describe 需要在数据库配置

key 搜索的关键字

key_enroll 是否开启关键字词库录取


收集词库后并不会时时更新，需要重启服务


```
get 请求
http://localhost:8084/zldz/trie/keysearch?trie_type.id=1&trie_type.trie_describe=测试字典树&key=搜索的提示关键词&key_enroll=true
```


# 具体接入步骤

- 联系管理员，申请一个trieType
- 提供词库录入（非必须）
- 调用接口  /zldz/trie/keysearch
- 需要更新字典树的，请联系管理员重启服务，管理员比较懒，没有写crontab


