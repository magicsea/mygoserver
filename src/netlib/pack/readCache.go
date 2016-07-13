package pack

import (
	binary "encoding/binary"
)

type ReadyCache struct {
	cbuff        *CircleBuffer
	completePack chan<- NetPack
	workSize     uint16
}

func NewReadyCache(completePack chan<- NetPack) *ReadyCache {
	cbuffer := &pack.CircleBuffer{}
	cbuffer.Ini(1024 * 64)
	return &ReadyCache{cbuffer, completePack, 0}
}

func (self *ReadyCache) AddData(data []byte) {

	self.cbuff.Write(data)
	ParsePack()
}

func (self *ReadyCache) ParsePack() {
	//read head ,2 size+2 code
	if self.workSize > 0 {
		//have read size,read other
		if self.cbuff.Size() >= self.workSize {
			data := make([]byte, self.workSize)
			readsize, ok = self.cbuff.Read(data)
			if ok {
				code := binary.BigEndian.Uint16(data[:2])
				content := data[2:]
				pack := &packNetPack{self.workSize, code, content}
				completePack <- pack //complete one
				if self.cbuff.Size() > 0 {
					self.ParsePack()
				}
			}
		}
	} else if self.cbuff.Size() > 2 {
		//reading size
		data := make([]byte, 2)
		readsize, ok = self.cbuff.Read(data)
		if ok {
			size := binary.BigEndian.Uint16(data)
			if self.cbuff.Size() >= self.workSize {
				self.ParsePack()
			}
		}
	}
}
