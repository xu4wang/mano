package Mo

import (
	"fmt"
	"log"
)

//Mo is the Object for all the Mo
type Mo struct {
}

func (d *Mo) Info(args ...interface{}) {
	log.Println("[INFO]", fmt.Sprint(args...))
}

func (d *Mo) Infof(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func (d *Mo) Warning(args ...interface{}) {
	log.Println("[WARN]", fmt.Sprint(args...))
}

func (d *Mo) Warningf(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}

func (d *Mo) Error(args ...interface{}) {
	log.Println("[ERROR]", fmt.Sprint(args...))
}

func (d *Mo) Errorf(format string, args ...interface{}) {
	log.Printf("[ERROR] "+format, args...)
}
