# MEDcury Test

Open Terminal or CMD
```bash
command $ go test => for unit test main_test.go
command $ go run main.go => for run main.go
command $ .\MEDcury_test.exe => for run main.go
```
* TEST NO.1
    * WE ARE THE BACKYARD MY FRIEND

* TEST NO.2
    * Phone => 1-541-754-3010, Name => J Steeve, Address => 156 Alphand St.
    * Phone => 1-741-984-3090, Name => Peter Reedgrave, Address => Chicago
    * Error => Too many people: 1-098-512-2222
    * Error => Not found: 5-555-555-5555

#### result unit test

Running tool: C:\Go\bin\go.exe test -timeout 30s -run **^TestStringDecoder$**

PASS

ok  	_/d_/Benz/MEDcury_test	0.283s

Running tool: C:\Go\bin\go.exe test -timeout 30s -run **^TestPhoneBooks$**

PASS

ok  	_/d_/Benz/MEDcury_test	0.302s
