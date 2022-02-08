package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//这里将这些方法关联到结构体
type Transfer struct {
	//分析他应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这时传输时，使用缓冲
}

func (this *Transfer) ReadPkg()(mes message.Message,err error) {
	//buffer := make([]byte, 8096)
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil{
		//fmt.Println("读取错误",err)
		return
	}
	// 根据buffer[:4] 转成uint32类型,获取到需要读取多少个字节数
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	// 根据pkgLen 读取内容,存放到buffer数组中
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil{

		// 自定义错误
		//err = errors.New("读取错误~")
		return
	}

	// 将buffer 反序列化成结构体,需要传结构体的地址!!!!
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil{

		//err = errors.New("反序列化错误")
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buffer [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}