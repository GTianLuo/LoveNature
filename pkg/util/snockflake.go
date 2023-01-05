package util

import (
	"fmt"
	"lovenature/log"
	"math/rand"
	"sync"
	"time"
)

type Snowflake struct {
	mu           sync.Mutex
	timeStamp    int64 //时间戳,毫秒
	workId       int64 //工作节点id
	dataCenterId int64 //中心机房id
	sequence     int64 //序列号
}

const (
	epoch = int64(1672905141503) // 2023 1 5

	timeStampBits    = uint(41)
	dataCenterIdBits = uint(2)
	workIdBits       = uint(3)
	sequenceBits     = uint(12)

	dataCenterIdMax = int64(-1 ^ (-1 << dataCenterIdBits))
	workIdMax       = int64(-1 ^ (-1 << workIdBits))
	sequenceMax     = int64(-1 ^ (-1 << sequenceBits))
	timeStampMax    = int64(-1 ^ (-1 << timeStampBits))

	workIdShift       = sequenceBits
	dataCenterIdShift = sequenceBits + workIdBits
	timeStampShift    = sequenceBits + workIdBits + dataCenterIdBits
)

var IdGenerator *Snowflake = &Snowflake{
	dataCenterId: 1,
	workId:       1,
	sequence:     1,
}

func (s *Snowflake) NextVal() int64 {
	s.mu.Lock()
	now := time.Now().UnixMilli()
	//时间戳相同
	if s.timeStamp == now {
		//序列号加一, 同时要保证序列号不能超出12位
		s.sequence = sequenceMax & (s.sequence + 1)
		//说明此时的序列号已经超出了范围，需要等待
		if s.sequence == 0 {
			for now <= s.timeStamp {
				now = time.Now().UnixMilli()
			}
			s.sequence = 1
		}
	} else {
		//时间不同，序列号从1开始
		s.sequence = 1
	}
	t := now - epoch
	//fmt.Println(t)
	if t > timeStampMax {
		s.mu.Unlock()
		log.Errorf("epoch : %v 已经不支持当前时间的id生成")
		return 0
	}
	r := int64((t << timeStampShift) | (s.dataCenterId << dataCenterIdShift) | (s.workId << workIdShift) | s.sequence)
	s.timeStamp = now
	s.mu.Unlock()
	return r
}

var m map[int]string

func init() {
	m = make(map[int]string, 10)
	m[0] = "红眼树蛙"
	m[1] = "小金毛"
	m[2] = "含羞草"
	m[3] = "捕蝇草"
	m[4] = "猪笼草"
	m[5] = "独角仙"
	m[6] = "小香猪"
	m[7] = "科尔鸭"
}

func GetNickName() string {
	return fmt.Sprintf("%s%x", m[rand.Intn(8)], IdGenerator.NextVal())
}
