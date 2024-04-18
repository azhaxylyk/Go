# ascii-art-fs

## Objectives

1) Ascii-art-fs is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII banners. 
2) This project should handle an input with numbers, letters, spaces, special characters and \n.
3) Your input must be in double quotes
4) You should conside output area
5) The usage must respect this format go run . [STRING] [BANNER], any other formats must return the following usage message:
```
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```
6) If there are other ascii-art optional projects implemented, the program should accept other correctly formatted [OPTION] and/or [BANNER].
Additionally, the program must still be able to run with a single [STRING] argument.

## Banner Format

- Each character has a height of 8 lines.
- Characters are separated by a new line \n.
- Bannernames should be standard/shadow/thinkertoy

## Usage
1) Clone the repository 
```
git clone git@git.01.alem.school:diyarulin/ascii-art-fs.git
```
2) Open folder with main.go file
```
cd main/
```
3) Run the program with following template

```
$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

### Unit Testing Usage

```
cd functional
go test -v
```

### Authors of the project

- [Didar (Diyarulin)](https://01.alem.school/git/diyarulin)
- [Aknur (Azhaxylyk)](https://01.alem.school/git/azhaxylyk)