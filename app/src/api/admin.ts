import requests from '@/server/request'
export const listDepartment =<T>(data:T)=> requests('/department/list',data,'get')
export const addDepartment = <T>(data:T) => requests('/department/add',data)
export const redactDepartment = <T>(data:T) => requests('/department/redact',data)
export const deleteDepartment = <T>(data:T) => requests('/department/delete',data)