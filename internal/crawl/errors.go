package crawl

import "errors"

var (
	errDuplicatedRecord = errors.New("duplicated record")
)

func isErrDuplicatedRecord(err error) bool {
	return errors.Is(err, errDuplicatedRecord)
}
