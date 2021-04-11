# Go (Golang) GOOS and GOARCH

All of the following information is based on `go version go1.14.7 darwin/amd64`.

## A list of valid GOOS values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- **`aix`**
- `android`
- **`darwin`**
- **`dragonfly`**
- **`freebsd`**
- `hurd`
- **`illumos`**
- **`js`**
- **`linux`**
- `nacl`
- **`netbsd`**
- **`openbsd`**
- **`plan9`**
- **`solaris`**
- **`windows`**
- `zos`

> NOTE: `nacl` support was dropped since `go version 1.14`.

## A list of valid GOARCH values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- **`386`**
- **`amd64`**
- `amd64p32`
- **`arm`**
- `armbe`
- **`arm64`**
- `arm64be`
- **`ppc64`**
- **`ppc64le`**
- **`mips`**
- **`mipsle`**
- **`mips64`**
- **`mips64le`**
- `mips64p32`
- `mips64p32le`
- `ppc`
- `riscv`
- **`riscv64`**
- `s390`
- **`s390x`**
- `sparc`
- `sparc64`
- **`wasm`**

> NOTE: The `amd64p32` GOARCH, which is related to the `nacl` GOOS, was dropped since `go version 1.14`.

## A list of valid 32-bit GOARCH values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- **`386`**
- `amd64p32`
- **`arm`**
- `armbe`
- **`mips`**
- **`mipsle`**
- `mips64p32`
- `mips64p32le`
- `ppc`
- `riscv`
- `s390`
- `sparc`

## A list of valid 64-bit GOARCH values

(Bold = supported by `go` out of the box, ie. without the help of a C compiler, etc.)

- **`amd64`**
- **`arm64`**
- `arm64be`
- **`ppc64`**
- **`ppc64le`**
- **`mips64`**
- **`mips64le`**
- **`riscv64`**
- **`s390x`**
- `sparc64`
- **`wasm`**

## A list of GOOS/GOARCH supported by `go` out of the box

- `aix/ppc64`
- `darwin/386`
- `darwin/amd64`
- `dragonfly/amd64`
- `freebsd/386`
- `freebsd/amd64`
- `freebsd/arm`
- `freebsd/arm64`
- `illumos/amd64`
- `js/wasm`
- `linux/386`
- `linux/amd64`
- `linux/arm`
- `linux/arm64`
- `linux/ppc64`
- `linux/ppc64le`
- `linux/mips`
- `linux/mipsle`
- `linux/mips64`
- `linux/mips64le`
- `linux/riscv64`
- `linux/s390x`
- `netbsd/386`
- `netbsd/amd64`
- `netbsd/arm`
- `netbsd/arm64`
- `openbsd/386`
- `openbsd/amd64`
- `openbsd/arm`
- `openbsd/arm64`
- `plan9/386`
- `plan9/amd64`
- `plan9/arm`
- `solaris/amd64`
- `windows/386`
- `windows/amd64`
- `windows/arm`

## A list of 32-bit GOOS/GOARCH supported by `go` out of the box

- `darwin/386`
- `freebsd/386`
- `freebsd/arm`
- `linux/386`
- `linux/arm`
- `linux/mips`
- `linux/mipsle`
- `netbsd/386`
- `netbsd/arm`
- `openbsd/386`
- `openbsd/arm`
- `plan9/386`
- `plan9/arm`
- `windows/386`
- `windows/arm`

## A list of 64-bit GOOS/GOARCH supported by `go` out of the box

- `aix/ppc64`
- `darwin/amd64`
- `dragonfly/amd64`
- `freebsd/amd64`
- `freebsd/arm64`
- `illumos/amd64`
- `js/wasm`
- `linux/amd64`
- `linux/arm64`
- `linux/ppc64`
- `linux/ppc64le`
- `linux/mips64`
- `linux/mips64le`
- `linux/riscv64`
- `linux/s390x`
- `netbsd/amd64`
- `netbsd/arm64`
- `openbsd/amd64`
- `openbsd/arm64`
- `plan9/amd64`
- `solaris/amd64`
- `windows/amd64`

## Support Grid 1

|                   | `android` | `darwin` | `js` | `linux` | `nacl` | `windows` |                   |
| ----------------: | :-------: | :------: | :--: | :-----: | :----: | :-------: | :---------------- |
| **`386`**         | `α`       | `O`      |      | `O`     |        | `O`       | **`386`**         |
| **`amd64`**       | `α`       | `O`      |      | `O`     |        | `O`       | **`amd64`**       |
| **`amd64p32`**    |           |          |      |         |        |           | **`amd64p32`**    |
| **`arm`**         | `α`       | `β1`     |      | `O`     |        | `O`       | **`arm`**         |
| **`armbe`**       |           |          |      |         |        |           | **`armbe`**       |
| **`arm64`**       | `α`       | `β2`     |      | `O`     |        |           | **`arm64`**       |
| **`arm64be`**     |           |          |      |         |        |           | **`arm64be`**     |
| **`ppc64`**       |           |          |      | `O`     |        |           | **`ppc64`**       |
| **`ppc64le`**     |           |          |      | `O`     |        |           | **`ppc64le`**     |
| **`mips`**        |           |          |      | `O`     |        |           | **`mips`**        |
| **`mipsle`**      |           |          |      | `O`     |        |           | **`mipsle`**      |
| **`mips64`**      |           |          |      | `O`     |        |           | **`mips64`**      |
| **`mips64le`**    |           |          |      | `O`     |        |           | **`mips64le`**    |
| **`mips64p32`**   |           |          |      |         |        |           | **`mips64p32`**   |
| **`mips64p32le`** |           |          |      |         |        |           | **`mips64p32le`** |
| **`ppc`**         |           |          |      |         |        |           | **`ppc`**         |
| **`riscv`**       |           |          |      |         |        |           | **`riscv`**       |
| **`riscv64`**     |           |          |      | `O`     |        |           | **`riscv64`**     |
| **`s390`**        |           |          |      |         |        |           | **`s390`**        |
| **`s390x`**       |           |          |      | `O`     |        |           | **`s390x`**       |
| **`sparc`**       |           |          |      |         |        |           | **`sparc`**       |
| **`sparc64`**     |           |          |      | `γ`     |        |           | **`sparc64`**     |
| **`wasm`**        |           |          | `O`  |         |        |           | **`wasm`**        |
| | **`android`** | **`darwin`** | **`js`** | **`linux`** | **`nacl`** | **`windows`** | |

> NOTE: `darwin` is essentially the same as **macOS** / **iOS** ( https://golang.org/doc/install/source ).

> NOTE: `nacl` support was dropped since `go version 1.14`.

## Support Grid 2

|                   | `a` | `d` | `f` | `h` | `i` | `n` | `o` | `p` | `s` | `z` |                   |
| ----------------: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :---------------- |
| **`386`**         |     |     | `O` |     |     | `O` | `O` | `O` |     |     | **`386`**         |
| **`amd64`**       |     | `O` | `O` |     | `O` | `O` | `O` | `O` | `O` |     | **`amd64`**       |
| **`amd64p32`**    |     |     |     |     |     |     |     |     |     |     | **`amd64p32`**    |
| **`arm`**         |     |     | `O` |     |     | `O` | `O` | `O` |     |     | **`arm`**         |
| **`armbe`**       |     |     |     |     |     |     |     |     |     |     | **`armbe`**       |
| **`arm64`**       |     |     | `O` |     |     | `O` | `O` |     |     |     | **`arm64`**       |
| **`arm64be`**     |     |     |     |     |     |     |     |     |     |     | **`arm64be`**     |
| **`ppc64`**       | `O` |     |     |     |     |     |     |     |     |     | **`ppc64`**       |
| **`ppc64le`**     |     |     |     |     |     |     |     |     |     |     | **`ppc64le`**     |
| **`mips`**        |     |     |     |     |     |     |     |     |     |     | **`mips`**        |
| **`mipsle`**      |     |     |     |     |     |     |     |     |     |     | **`mipsle`**      |
| **`mips64`**      |     |     |     |     |     |     |     |     |     |     | **`mips64`**      |
| **`mips64le`**    |     |     |     |     |     |     |     |     |     |     | **`mips64le`**    |
| **`mips64p32`**   |     |     |     |     |     |     |     |     |     |     | **`mips64p32`**   |
| **`mips64p32le`** |     |     |     |     |     |     |     |     |     |     | **`mips64p32le`** |
| **`ppc`**         |     |     |     |     |     |     |     |     |     |     | **`ppc`**         |
| **`riscv`**       |     |     |     |     |     |     |     |     |     |     | **`riscv`**       |
| **`riscv64`**     |     |     |     |     |     |     |     |     |     |     | **`riscv64`**     |
| **`s390`**        |     |     |     |     |     |     |     |     |     |     | **`s390`**        |
| **`s390x`**       |     |     |     |     |     |     |     |     |     |     | **`s390x`**       |
| **`sparc`**       |     |     |     |     |     |     |     |     |     |     | **`sparc`**       |
| **`sparc64`**     |     |     |     |     |     |     |     |     |     |     | **`sparc64`**     |
| **`wasm`**        |     |     |     |     |     |     |     |     |     |     | **`wasm`**        |
| | **`a`** | **`d`** | **`f`** | **`h`** | **`i`** | **`n`** | **`o`** | **`p`** | **`s`** | **`z`** |

### Key

`a` = `aix`, `d` = `dragonfly`, `f` = `freebsd`, `h` = `hurd`, `i` = `illumos`,
`n` = `netbsd`, `o` = `openbsd`, `p` = `plan9`, `s` = `solaris`, `z` = `zos`

(blank): Unsupported

`O`: Supported

`α`:
```
# command-line-arguments
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: unknown option: -z
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`β1`:
```
# command-line-arguments
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: warning: ignoring file /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o, building for macOS-x86_64 but attempting to link with file built for unknown-armv7
Undefined symbols for architecture x86_64:
  "_main", referenced from:
     implicit entry/start for main executable
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`β2`:
```
# command-line-arguments
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: warning: ignoring file /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxx/go.o, building for macOS-x86_64 but attempting to link with file built for unknown-arm64
Undefined symbols for architecture x86_64:
  "_main", referenced from:
     implicit entry/start for main executable
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`γ`:
```
go tool compile: exit status 2
compile: unknown architecture "sparc64"
go tool compile: exit status 2
compile: unknown architecture "sparc64"
```

## Source 1

### main.go

```go
package main

func main() {}
```

### make.sh

```sh
#!/bin/sh

os_archs=()

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
for goos in aix android darwin dragonfly freebsd hurd illumos js \
            linux nacl netbsd openbsd plan9 solaris windows zos
do
    for goarch in 386 amd64 amd64p32 arm armbe arm64 arm64be ppc64 \
                  ppc64le mips mipsle mips64 mips64le mips64p32 \
                  mips64p32le ppc riscv riscv64 s390 s390x sparc sparc64 wasm
    do
        echo "--------"
        echo "${goos}/${goarch}"
        echo "--------"
        GOOS=${goos} GOARCH=${goarch} go build -o /dev/null main.go >/dev/null 2>&1
        if [ $? -eq 0 ]
        then
            os_archs+=("${goos}/${goarch}")
        fi
    done
done

echo
echo "================"
echo
for os_arch in "${os_archs[@]}"
do
    echo ${os_arch}
done
```

## Source 2

### main.go

```go
package main

const (
	hello uint = 0xfedcba9876543210
)

func main() {}
```

### make.sh

```sh
#!/bin/bash

# Reference:
# https://github.com/golang/go/blob/master/src/go/build/syslist.go
os_archs=(
    aix/ppc64
    darwin/386
    darwin/amd64
    dragonfly/amd64
    freebsd/386
    freebsd/amd64
    freebsd/arm
    freebsd/arm64
    illumos/amd64
    js/wasm
    linux/386
    linux/amd64
    linux/arm
    linux/arm64
    linux/ppc64
    linux/ppc64le
    linux/mips
    linux/mipsle
    linux/mips64
    linux/mips64le
    linux/riscv64
    linux/s390x
    netbsd/386
    netbsd/amd64
    netbsd/arm
    netbsd/arm64
    openbsd/386
    openbsd/amd64
    openbsd/arm
    openbsd/arm64
    plan9/386
    plan9/amd64
    plan9/arm
    solaris/amd64
    windows/386
    windows/amd64
    windows/arm
)

os_archs_32=()
os_archs_64=()

for os_arch in "${os_archs[@]}"
do
    goos=${os_arch%/*}
    goarch=${os_arch#*/}
    GOOS=${goos} GOARCH=${goarch} go build -o /dev/null main.go >/dev/null 2>&1
    if [ $? -eq 0 ]
    then
        os_archs_64+=(${os_arch})
    else
        os_archs_32+=(${os_arch})
    fi
done

echo "32-bit:"
for os_arch in "${os_archs_32[@]}"
do
    printf "\t%s\n" "${os_arch}"
done
echo

echo "64-bit:"
for os_arch in "${os_archs_64[@]}"
do
    printf "\t%s\n" "${os_arch}"
done
echo
```