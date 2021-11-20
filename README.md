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

```
$ brew tap webcyou-org/coppeno
$ brew install coppeno
```

**golang**

```
go install github.com/webcyou-org/coppeno@latest
```

or

```
GO111MODULE=on go get github.com/webcyou-org/coppeno
```

## Usage Example

**Default Usage**

```
$ coppeno <command>
```

#### Commands

| Commands | argument | description    |
|:----------------|:-----------|:-----------|
| init            | None |  Initialize coppeno.json. |
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

```
NAME:
   coppeno - Quick project kickstarter Simple File Manager CLI tool.

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.7.2

COMMANDS:
   init      Initialize coppeno.json.
   save, s   Save the file name and file path interactively.
   load, l   Read the coppeno.json file that was created.
   fetch, f  Download the file saved in coppeno.json
   list      Check the list of files saved in coppeno.json.
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Quick Start

#### dotfiles

**step1**

```
coppeno load https://github.com/webcyou-org/coppeno-json/blob/master/dotfiles.json
```

**step2**

```
coppeno fetch
```


#### cats

**step1**

```
coppeno load https://github.com/webcyou-org/coppeno-json/blob/master/cats.json
```

**step2**

```
coppeno fetch
```

## config files

 Create a file that conforms to the [XDG Base Directory specification](https://specifications.freedesktop.org/basedir-spec/latest/).

### coppeno.json

```
{
  // Not grouped.
  "none": [
    {
      "name": "File Name",
      "url": "File URL for download"    
    },
    {
      "name": "File Name",
      "url": "File URL for download"    
    }
  ],
  // Arbitrarily grouped.
  "dotfiles": [
    {
      "name": ".gitconfig",
      "url": "https://github.com/lewagon/dotfiles/blob/master/gitconfig"
    },
    {
      "name": ".bash_profile",
      "url": "https://github.com/mathiasbynens/dotfiles/blob/main/.bash_profile"
    }
  ]
}
```

### config.json

```
```

### coppeno struct

```
type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
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