# go-featrues
go featrues 

## efficiently go  
### string join 
0. '+' (base):字符串越长越慢 
1. string.Builder(>v1.10) (~5000): https://golang.org/pkg/strings/#Builder 
2. bytes.Buffer(~2000) : https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go 
[benchmark](https://github.com/craftsdong/go-featrues/blob/master/string_test.go) 

### capacity 
1. map ,slice 在能预知容量的情况下指定cap，可以有效减少动态伸缩和GC 
[benchmark](https://github.com/craftsdong/go-featrues/blob/master/capacity_test.go)

### panic 
1. [handle](https://golang.org/ref/spec#Handling_panics) 
2. [running](https://golang.org/ref/spec#Run_time_panics)
### recover in goroutniue
持续运行的守候型服务，需要考虑pinic-recover情况，以免服务运行越来越少，最后可能无服务可用
1. [错误示范](https://github.com/craftsdong/go-featrues/blob/master/goroutniue_incorrect.go) 
2. [正确示范](https://github.com/craftsdong/go-featrues/blob/master/goroutniue_correct.go) 
3. [正确示范结构](https://github.com/craftsdong/continuous_daemon)  
