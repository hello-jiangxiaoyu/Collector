package output

import "Collector/output/database"

type UploadLog interface {
	SendLog(map[string]any) error
}

func GetStore(t string) (UploadLog, error) {
	return database.NewGormOutput(t, "")
}

func FlushLog(itf UploadLog, log map[string]any) error {
	return itf.SendLog(log)
}
