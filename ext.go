package queue

// #include <stdlib.h>
// #include "queue.h"
import "C"
import (
	"unsafe"

	"github.com/dunglas/frankenphp"
)

func init() {
	frankenphp.RegisterExtension(unsafe.Pointer(&C.queue_module_entry))
}

//export frankenphp_queue
func frankenphp_queue(data *C.zval) {
	queue(data)
}
