package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

// create Hash Object
func Instance() *Hash {
	return &Hash{}
}

type Hash struct {
}

//md5哈希
func (this *Hash) Md5(str string) (md5Str string) {
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	md5Str = fmt.Sprintf("%x", md5Inst.Sum([]byte("")))
	return
}

//sha1哈希
func (this *Hash) Sha1(str string) (sha1Str string) {
	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(str))
	sha1Str = fmt.Sprintf("%x", sha1Inst.Sum([]byte("")))
	return
}

//md5文件哈希
func (this *Hash) Md5File() {

}

//sha1文件哈希
func (this *Hash) Sha1File() {

}

// BUG (Tabalt) : #1 md5文件哈希 实现
// BUG (Tabalt) : #2 sha1文件哈希 实现
