# go-featrues
go featrues 

## efficiently go  
### string join 
0. '+' (base):字符串越长越慢 
1. string.Builder (~5000): https://golang.org/pkg/strings/#Builder 
2. bytes.Buffer(~2000) : https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go 
![beachmark](https://github.com/craftsdong/go-framework/blob/master/string_join_beach.png) 

### capacity 
1. map ,slice 在能预知容量的情况下指定cap，可以有效减少动态伸缩和GC
