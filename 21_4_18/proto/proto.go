package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {

	//读取信息的长度，转化成int32类型(占4字节)
	length := int32(len(message))
	pkg := new(bytes.Buffer)

	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil

}

func Decode(reader *bufio.Reader) (string, error) {

	//读取消息头
	msgHeadBuf, _ := reader.Peek(4) //读取前4个字节的数据
	lengthBuf := bytes.NewBuffer(msgHeadBuf)

	var length int32
	err := binary.Read(lengthBuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil

}
