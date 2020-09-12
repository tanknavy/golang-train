package main
import "fmt"

//https://www.davidkaya.com/sets-in-golang/
//https://github.com/deckarep/golang-set/blob/master/threadunsafe.go
//可以使用以下自定义map实现set
//为什么key用interface{}而value用struct{}, 涉及到类型使用interface{},涉及值使用struct{}{}
var InterfaceSet map[interface{}]struct{}
var IntSet       map[int]struct{} 
var StringSet    map[string]struct{}

var InterfaceSet2 map[interface{}]bool //bool为true表示有这个key就存在
var IntSet2       map[int]bool
var StringSet2    map[string]bool

var exists = struct{}{} //没有字段的struct

type set struct { //使用这个定义set类型
    m map[string]struct{}
}

func NewSet() *set {
    s := &set{}
    s.m = make(map[string]struct{})
    return s
}

func (s *set) Add(value string) {
    s.m[value] = exists
}

func (s *set) Remove(value string) {
    delete(s.m, value)
}

func (s *set) Contains(value string) bool {
    _, c := s.m[value] //不存在的key返回0
    return c
}

func main() {
    s := NewSet()

    s.Add("Peter")
    s.Add("David")

    fmt.Println(s.Contains("Peter"))  // True
    fmt.Println(s.Contains("George")) // False

    s.Remove("David")
    fmt.Println(s.Contains("David")) // False
}