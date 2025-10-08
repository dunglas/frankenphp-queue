#include <Zend/zend_API.h>
#include <Zend/zend_hash.h>
#include <Zend/zend_types.h>
#include <php.h>
#include <stddef.h>

#include "_cgo_export.h"
#include "queue.h"
#include "queue_arginfo.h"

PHP_MINIT_FUNCTION(queue) { return SUCCESS; }

zend_module_entry queue_module_entry = {STANDARD_MODULE_HEADER,
                                        "queue",
                                        ext_functions,    /* Functions */
                                        PHP_MINIT(queue), /* MINIT */
                                        NULL,             /* MSHUTDOWN */
                                        NULL,             /* RINIT */
                                        NULL,             /* RSHUTDOWN */
                                        NULL,             /* MINFO */
                                        "1.0.0",          /* Version */
                                        STANDARD_MODULE_PROPERTIES};

PHP_FUNCTION(frankenphp_queue) {
  zval *data = NULL;

  ZEND_PARSE_PARAMETERS_START(1, 1)
  Z_PARAM_ZVAL(data)
  ZEND_PARSE_PARAMETERS_END();

  frankenphp_queue(data);
}
