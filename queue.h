#ifndef _QUEUE_H
#define _QUEUE_H

#include <php.h>
#include <stdint.h>

extern zend_module_entry queue_module_entry;

static void zval_copy_value(zval *dst, zval *src) { ZVAL_COPY_VALUE(dst, src); }

#endif
