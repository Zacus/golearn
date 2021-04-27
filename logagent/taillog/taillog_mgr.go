package taillog

import "golearn/logagent/etcd"

var tskMgr *taillogMgr

type taillogMgr struct {
	LogEntry []*etcd.LogEntry
	// tskMap   map[string]*TailTask
}

func Init(LogEntryConf []*etcd.LogEntry) {

	tskMgr = &taillogMgr{
		LogEntry: LogEntryConf, //保存当前日志收集配置
	}
	for _, LogEntry := range LogEntryConf {

		NewTailTask(LogEntry.Path, LogEntry.Topic)

	}
}
