import dictApi from './dictIndex'
import cronApi from './cron'
import taskApi from './task'

export const getDictListApi = dictApi.getDictListApi
export const addDictApi = dictApi.addDictApi
export const editDictApi = dictApi.editDictApi
export const delDictApi = dictApi.delDictApi
export const getDictDataListApi = dictApi.getDictDataListApi
export const addDictDataApi = dictApi.addDictDataApi
export const editDictDataApi = dictApi.editDictDataApi
export const delDictDataApi = dictApi.delDictDataApi

export const getCronListApi = cronApi.getCronListApi
export const addCronApi = cronApi.addCronApi
export const editCronApi = cronApi.editCronApi
export const delCronApi = cronApi.delCronApi
export const editCronStatusApi = cronApi.editCronStatusApi

export const getTaskListApi = taskApi.getTaskListApi
export const addTaskApi = taskApi.addTaskApi
export const delTaskApi = taskApi.delTaskApi
