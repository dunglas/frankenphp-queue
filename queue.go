package queue

// #include <Zend/zend_types.h>
// #include "queue.h"
import "C"
import (
	"log/slog"
	"unsafe"
)

//export_php:function frankenphp_queue(mixed $data): void
func queue(data *C.zval) {
	var p C.zval
	C.zval_copy_value(&p, data)
	defer C.zval_ptr_dtor(&p)

	ret := worker.SendMessage(unsafe.Pointer(&p), nil)
	logger.Debug("frankenphp_queue() finished", slog.Any("parameters", p), slog.Any("return", ret))
}
