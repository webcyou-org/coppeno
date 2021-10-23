<div align="center">
  <h1><code>coppeno</code></h1>
  <p>
    <strong>Quick project kickstarter Simple File Manager CLI tool.</strong>
  </p>
</div>

![linkedin_banner_image_1](https://user-images.githubusercontent.com/1584153/138566408-3fd46669-c7be-42ce-afcc-eaf52bf6f0ff.png)

![IMG_0280](https://user-images.githubusercontent.com/1584153/118537557-0eff9c00-b788-11eb-96cb-6caaf76c3835.JPG)

![coppeno-3](https://user-images.githubusercontent.com/1584153/118579050-468e3880-b7c8-11eb-9365-98addbd8a363.jpg)


## Installation

go

```
$ GO111MODULE=on go get github.com/webcyou-org/coppeno
```

## Usage Example

**Default Usage**



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

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

## Creators

**Daisuke Takayama**
* [@webcyou](https://twitter.com/webcyou)
* [@panicdragon](https://twitter.com/panicdragon)
* <https://github.com/webcyou>
* <https://github.com/panicdragon>
* <http://www.webcyou.com/>

## Copyright and license
MIT