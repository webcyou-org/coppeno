<div align="center">
  <h1><code>coppeno</code></h1>
  <p>
    <strong>Quick project kickstarter Simple File Manager CLI tool.</strong>
  </p>
</div>

![linkedin_banner_image_1](https://user-images.githubusercontent.com/1584153/138566408-3fd46669-c7be-42ce-afcc-eaf52bf6f0ff.png)

## Usage Image

![coppeno](https://user-images.githubusercontent.com/1584153/138567391-f206a6bc-7834-4157-bd8c-3dedf415eff4.png)

## Installation

**Homebrew**

**golang**

```
$ GO111MODULE=on go get github.com/webcyou-org/coppeno
```

## Usage Example

**Default Usage**

```
$ coppeno <command>
```

#### Commands

| Commands | argument | description    |
|:----------------|:-----------|:-----------|
| save            | [filename] [URL] |  Save the file name and file path interactively. |
| load            | [filepath] |  Read the coppeno.json file that was created. |
| fetch           | None or [filename] |  Download the file saved in coppeno.json |
| diff            |  |  Check the difference between the destination file and the local file. |
| list            | None |  Check the list of files saved in coppeno.json. |
| setting         | None |  Updating the configuration file |
| cache           |  |  Cache the downloaded files |

To see a full list of commands available for the CLI tools, type:

```
$ coppeno help
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