package queue

// #include <Zend/zend_types.h>
// #include "queue.h"
import "C"
import (
	"unsafe"

	"github.com/dunglas/frankenphp"
	"go.uber.org/zap"
)

//export_php:function frankenphp_queue(mixed $data): void
func queue(data *C.zval) {
	var p C.zval
	C.zval_copy_value(&p, data)

	w.InjectRequest(&frankenphp.WorkerRequest{
		CallbackParameters: unsafe.Pointer(&p),
		// Less efficient, but simpler alternative (do more operations)
		//CallbackParameters:	frankenphp.GoValue(unsafe.Pointer(data)),
		AfterFunc: func(callbackReturn any) {
			C.zval_ptr_dtor(&p)
			w.logger.Debug("frankenphp_queue() finished", zap.Any("parameters", p), zap.Any("return", callbackReturn))
		},
	})
}
