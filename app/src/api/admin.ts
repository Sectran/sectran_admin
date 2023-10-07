import requests from '@/server/request'
export const listDepartment =<T>(data:T)=> requests('/department/list',data,'get')
export const addDepartment = <T>(data:T) => requests('/department/add',data)
export const editDepartment = <T>(data:T) => requests('/department/edit',data)
export const deleteDepartment = <T>(data:T) => requests('/department/delete',data) 


export const adduser = <T>(data:T) => requests('/user/add',data)