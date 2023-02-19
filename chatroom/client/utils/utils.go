package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"liaotian/chatroom/client/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//发送数据前需要整理数据
	// 1将data的长度发送给服务器
	// 2需要先获取data的长度 再转换成byte切片
	var pkglen uint32
	pkglen = uint32(len(data)) //转成无符号整形
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	//发送字节长度
	n, err := this.Conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("数据发送异常", err)
	}
	fmt.Println("发送消息数据长度:", len(data), "发送内容:", string(data))
	//再次发送字节内容
	_, err = this.Conn.Write(data)
	if err != nil {
		fmt.Println("发送数据异常", err)
		return
	}

	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("链接数据异常")
		return
	}
	//fmt.Println("读取长度:", buf[0:n])
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	n, err := this.Conn.Read(this.Buf[:pkgLen])

	if n != int(pkgLen) || err != nil {
		//err = errors.New("链接数据异常")
		return
	}
	//将 pkgLen反序列化->message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //需要传地址 使用指针
	if err != nil {
		//err = errors.New("反序列化失败")
		return
	}
	return mes, err
}
