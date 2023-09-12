package tail

import (
	"Collector/conf"
	"Collector/output"
	"encoding/json"
	"fmt"
	"github.com/hpcloud/tail"
	"strings"
	"time"
)

var (
	TailObj *tail.Tail
)

func Init() (err error) {
	fileName := conf.GetTailPath()
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	if TailObj, err = tail.TailFile(fileName, cfg); err != nil {
		fmt.Printf("tailf %s failed, err:%v\n", fileName, err)
		return err
	}

	return nil
}

func Run() error {
	itf, err := output.GetStore("")
	if err != nil {
		return err
	}

	for {
		line, ok := <-TailObj.Lines
		if !ok {
			fmt.Println("Tail file close reopen, filename", TailObj.Filename)
			time.Sleep(time.Second)
			continue // 重新打开日志文件
		}

		if len(strings.Trim(line.Text, "\r")) == 0 {
			continue // 过滤掉空行
		}

		logLine := make(map[string]any)
		if err = json.Unmarshal([]byte(line.Text), &logLine); err != nil {
			fmt.Println("unmarshal err", err)
			continue
		}

		if err = output.FlushLog(itf, logLine); err != nil {
			fmt.Println("send log err", err)
			continue
		}
	}
}
