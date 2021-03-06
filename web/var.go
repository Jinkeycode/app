package web

import "../ctsFrame/cacheTools"

type AdminAccountStruct struct {
	Username string
	Password string
}

type OwnConfigInfo struct {
	TestImage string
	HttpAddr string
	RedisObject cacheTools.RedisClientObject
	ImageMap map[string]string
	AdminAccount AdminAccountStruct
}

// REDIS KEY
var REDIS_KEY_APP_ZSET = "paas_app_zset"
var REDIS_KEY_APP_INFO_HSET = "paas_app_info_hset_"
var REDIS_KEY_APP_LOG_LIST = "paas_app_log_list_"

var REDIS_KEY_APP_DOMAIN_HSET = "paas_app_domain"
var REDIS_KEY_DOMAIN_APP_HSET = "paas_domain_app"

var REDIS_KEY_POST_USE = "paas_port_use_set"

var REDIS_KEY_APP_IMAGE_LIST = "paas_app_image_list_"

var REDIS_KEY_BUILD_IMAGE_TASK_LIST = "paas_build_image_task_list"

var REDIS_KEY_SSH_STR = "paas_ssh_str"
var REDIS_KEY_ERROR_SHELL_LIST = "paas_error_shell_list"
