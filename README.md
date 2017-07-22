dll_caller
==========

A windows dll and memory dll call hellper


## Windows MessageBox Example

```go
package main

import (
    "github.com/chennqqi/godll"
    "fmt"
)

func main(){
    ShowMessageBox()
}

func ShowMessageBox() {
    var dll *godll.Dll
    if d, e := godll.NewFileDll("user32.dll"); e != nil {
        fmt.Println(e.Error())
        return
    } else {
        dll = d
    }
	defer dll.FreeLibrary()

    if e := dll.InitalFunctions("MessageBoxW"); e != nil {
        fmt.Println(e.Error())
        return
    }

    ret, err := dll.Call("MessageBoxW", 0, "hello", "world", 3)

    fmt.Println(ret, err)
}
```

## Windows MessageBox Example by memdll

You can load dll for file/memory build with [go-bindata](github.com/jteeuwen/go-bindata)

```go
package main

import (
    "github.com/chennqqi/godll"
    "fmt"
	"io/ioutil"
)

func main(){
    ShowMessageBox()
}

func ShowMessageBox() {
    var dll *godll.MemDll

	dllbytes, _ := ioutil.ReadFile("user32.dll")
	//or you can load other 
	
    if d, e := godll.NewMemDll(dllbytes); e != nil {
        fmt.Println(e.Error())
        return
    } else {
        dll = d
    }
	defer dll.FreeLibrary()

    if e := dll.InitalFunctions("MessageBoxW"); e != nil {
        fmt.Println(e.Error())
        return
    }

    ret, err := dll.Call("MessageBoxW", 0, "hello", "world", 3)

    fmt.Println(ret, err)
}
```


## OTHER

	[x] win7 386 test ok
	[ ] win7 amd64 not tested

	