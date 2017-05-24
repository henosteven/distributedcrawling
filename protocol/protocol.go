package protocol

import (
    "bytes"
    "encoding/binary"
)

const (
    ConstHeader = "headerflag"
    ConstHeaderLen = 10
    ConstSaveDataLength = 4
)

func Pack(msg []byte) []byte{
    return append(append([]byte(ConstHeader), IntToByte(len(msg))...), msg...)
}

func UnPack(content []byte, ch chan []byte) []byte{

    var length int = len(content)

    var i int = 0; 
    for i =0; i < length; i++ {
        /* 长度不够包头 */
        if length < i+ ConstHeaderLen + ConstSaveDataLength {
            break
        }

        if string(content[i:i+ConstHeaderLen]) == ConstHeader {
            
            realLen := ByteToInt(content[i+ConstHeaderLen: i+ConstHeaderLen+ConstSaveDataLength])
            
            /* 包不完整 */
            if i+ConstHeaderLen+ConstSaveDataLength+realLen > length {
                break
            }
            
            /* 完整的数据就直接返回 */
            ch <- content[i+ConstHeaderLen+ConstSaveDataLength: i+ConstHeaderLen+ConstSaveDataLength+realLen]
            
            /* 因为上部分会有+1操作，所以这里提前加一处理 */
            i = i+ConstHeaderLen+ConstSaveDataLength+realLen - 1
        }

    }
    
    return content[i:]
}

func IntToByte(n int) []byte {
   x := int32(n) 
   bytesBuffer := bytes.NewBuffer([]byte{})
   binary.Write(bytesBuffer, binary.BigEndian, x)
   b := bytesBuffer.Bytes()
   return b
}

func ByteToInt(b []byte) int {
    bytesBuffer := bytes.NewBuffer(b)
    var x int32
    binary.Read(bytesBuffer, binary.BigEndian, &x)
    return int(x)
}
