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

## Quick Start

### dotfiles example

This is an example of downloading a dotfile.

**step1**

```
coppeno load https://github.com/webcyou-org/coppeno-json/blob/master/dotfiles.json
```

**step2**

```
coppeno fetch
```

Selecting "All" will download all the files you have set.

![dotfiles1](https://user-images.githubusercontent.com/1584153/142711100-de24005a-9989-42d3-a088-4e9ac1e114d7.png)

![dotfiles2](https://user-images.githubusercontent.com/1584153/142711110-5715a51b-cfe1-42bf-b974-946199f33521.png)

![dotfiles3](https://user-images.githubusercontent.com/1584153/142711112-98de9c16-2d70-4a9b-8d91-66b615e1c4cc.png)

### cats example

Here's an example of a cute cat.

**step1**

```
coppeno load https://github.com/webcyou-org/coppeno-json/blob/master/cats.json
```

**step2**

```
coppeno fetch
```

Selecting "All" will download all the cute cat images you have set.

![dotfiles1](https://user-images.githubusercontent.com/1584153/142711100-de24005a-9989-42d3-a088-4e9ac1e114d7.png)

![cat1](https://user-images.githubusercontent.com/1584153/142711380-9a6ad1ad-7103-4aad-823c-14fabe14dc12.png)

![cat2](https://user-images.githubusercontent.com/1584153/142711381-c58bcf49-53fa-46d7-8e75-3e4d2259248c.png)

Oh. This is kawaii.

![cat_all](https://user-images.githubusercontent.com/1584153/142711670-ed12cda0-5a62-41c4-ad82-79183e53cae8.png)


You can also load coppeno.json as a local file.

```
coppeno load ./sample/example.json
```

### Set any file

It is possible to register individual download files.

**step1**

```
coppeno save
```

Set the group name, file name, and download URL to be managed.The group name is optional.

![save1](https://user-images.githubusercontent.com/1584153/142712397-7743b36f-b2c3-4920-925f-ae49d80caaba.png)

![save2](https://user-images.githubusercontent.com/1584153/142712401-2def731c-efd5-4e9a-abef-0be6d38bd419.png)

**step2**

```
coppeno fetch
```

Download the files registered with `coppeno fetch`.

![save3](https://user-images.githubusercontent.com/1584153/142712396-c1a96cbb-f2bd-4fba-a06a-864474eaea58.png)

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
coming soon.
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