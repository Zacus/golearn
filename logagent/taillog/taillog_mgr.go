package taillog

import (
	"fmt"
	"golearn/logagent/etcd"
	"time"
)

var tskMgr *taillogMgr

//tailtsk 管理者
type taillogMgr struct {
	LogEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(LogEntryConf []*etcd.LogEntry) {

	tskMgr = &taillogMgr{
		LogEntry:    LogEntryConf, //保存当前日志收集配置
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //	无缓冲区通道
	}
	for _, LogEntry := range LogEntryConf {
		//初始化时记录tailtask个数,后续方便判断
		tailObj := NewTailTask(LogEntry.Path, LogEntry.Topic)
		mk := fmt.Sprintf("%s_%s", LogEntry.Path, LogEntry.Topic)
		tskMgr.tskMap[mk] = tailObj

	}
	go tskMgr.run()
}

//监听自己的newConfChan,有了新的配置过来之后就做对应的处理
//1.配置新增
//2.配置删除
//3.配置变更
func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:

			for _, conf := range newConf {

				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					continue
				} else {
					//新增配置
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}

			//删除newConf中没有,原来配置集中有的配置项
			for _, rawConf := range t.LogEntry {
				isDelete := true
				for _, nwConf := range newConf {
					if rawConf.Path == nwConf.Path && rawConf.Topic == nwConf.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					//停掉rawConf
					mk := fmt.Sprintf("%s_%s", rawConf.Path, rawConf.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
			fmt.Println("新的配置更新", newConf)

		default:
			time.Sleep(time.Second)
		}
	}

}

//外部暴露tskMgr newConfChan
func PullNewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
