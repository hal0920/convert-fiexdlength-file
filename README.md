# cvfv

CLI Tools for convert FixedLength file into VariableLength file

## Install

Use go get :

```sh
go get -u github.com/hal0920/cvfv
cd $GOPATH/src/github.com/hal0920/cvfv
go install
```

When using it please add the following to PATH :

```sh
export PATH=$PATH:$GOPATH/bin
```

## Setting

To use it please create the following setting file.

Unix : `~/.config/cvfv/settings.toml`

Write following :

```toml
[layout]
[layout.example1]
length =[1,2,3,4]

[layout.example2]
length =[4,3,2,1]
```

The length specification is the number of characters.
It is not the number of bytes.

## Usage

```text
NAME:
   cvfv - Convert Fixed-length file into variable-length file

USAGE:
   cvfv [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --layout FILELAYOUT, -l FILELAYOUT  Fixed-length file layout FILELAYOUT
   --help, -h                          show help (default: false)
   --version, -v                       print the version (default: false)
```

### Example of use

#### From File

```sh
cat test/test1.dat
abbcccdddd
1223334444

cvfv -l example1 test/test1.dat
a,bb,ccc,dddd
1,22,333,4444
```

#### From Standard input

```sh
echo 1223334444 | cvfv -l example1
1,22,333,4444
```
