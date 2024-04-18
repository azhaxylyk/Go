# ASCII Art Reverse

## Objectives

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. You will have to create a text file containing a graphic representation of a random string given as an argument.

The argument will be a flag, --reverse=<fileName>, in which --reverse is the flag and <fileName> is the file name. The program must then print this string in normal text.

## Getting Started

Clone the repository:
```bash
git clone git@git.01.alem.school:abulatov/ascii-art-reverse.git
```

Navigate to the project directory:
```bash
cd ascii-art-reverse
```

## Usage

```console
$ cat file.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
$ go run . --reverse=file.txt
hello
$
```

### authors
[abulatov](https://01.alem.school/git/abulatov)
<br>
[azhaxylyk](https://01.alem.school/git/azhaxylyk)