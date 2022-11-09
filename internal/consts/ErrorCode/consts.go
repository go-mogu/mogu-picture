package ErrorCode

const (
	QUERY_DEFAULT_ERROR            = 100
	SYSTEM_CONFIG_IS_NOT_EXIST     = 101
	PLEASE_CONFIGURE_PASSWORD      = 102
	PLEASE_CONFIGURE_BLOG_COUNT    = 103
	PLEASE_CONFIGURE_TAG_COUNT     = 104
	PLEASE_CONFIGURE_LINK_COUNT    = 105
	PLEASE_CONFIGURE_SYSTEM_PARAMS = 106
	PLEASE_SET_QI_NIU              = 107
	PLEASE_SET_LOCAL               = 108
	PLEASE_SET_MINIO               = 109
	SYSTEM_CONFIG_NOT_EXIST        = 110
	INSERT_DEFAULT_ERROR           = 200
	UPDATE_DEFAULT_ERROR           = 300
	DELETE_DEFAULT_ERROR           = 400
	DELETE_FAILED_PLEASE_CHECK_UID = 401
	OPERATION_SUCCESS              = 0
	OPERATION_FAIL                 = 1
	DATA_NO_EXIST                  = 2
	LOGIN_NOT_ACTIVE               = 3
	LOGIN_DISABLE                  = 4
	LOGIN_ERROR                    = 5
	INTERFACE_FREQUENTLY           = 6
	PARAM_INCORRECT                = 7
	INVALID_TOKEN                  = 8
	DATA_NO_PRIVILEGE              = 997
	ACCESS_NO_PRIVILEGE            = 998
	LOGIN_TIMEOUT                  = 999
	ERROR                          = "error"
)
