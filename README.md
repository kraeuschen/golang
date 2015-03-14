# sources
- https://github.com/nathany/vagrant-gopher/
- http://golang.org/doc/code.html
- https://gobyexample.com

# usage
```
vagrant up linux
vagrant ssh linux
cd /vagrant
export GOPATH=/vagrant/gocode/
export PATH=$PATH:$GOPATH/bin
```

# example
```
cd vagrant/gocode/src/github.com/kraeuschen/golang/basics/values
vi values.go
go install
values
```
