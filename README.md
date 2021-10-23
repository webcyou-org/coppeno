# coppeno

Quick project kickstarter Simple File Manager CLI tool.

![IMG_0280](https://user-images.githubusercontent.com/1584153/118537557-0eff9c00-b788-11eb-96cb-6caaf76c3835.JPG)

![coppeno-3](https://user-images.githubusercontent.com/1584153/118579050-468e3880-b7c8-11eb-9365-98addbd8a363.jpg)


```
$ coppeno コマンド
```

コマンド

```
save
load
update
diff
list
setting
cache
```



### coppeno.json （コンフィグファイル）

```
{
  # グループ
  nuxt: {
      key: 対象URL,
      key: 対象URL,
      key: 対象URL
  },
  〇〇: {
      key: 対象URL,
      key: 対象URL,
      key: 対象URL
  }
}
```

### coppeno struct

```
coppeno {
  filename,
  ext,
  domain,
  path,
  filesize
}
```

