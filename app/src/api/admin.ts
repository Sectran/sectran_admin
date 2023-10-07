import requests from '@/server/request'
export const listDepartment =<T>(data:T)=> requests('/department/list',data,'get')
export const addDepartment = <T>(data:T) => requests('/department/add',data)
export const editDepartment = <T>(data:T) => requests('/department/edit',data)
export const deleteDepartment = <T>(data:T) => requests('/department/delete',data) 


export const listUser =<T>(data:T)=> requests('/user/list',data,'get')
export const adduser = <T>(data:T) => requests('/user/add',data)
export const edituser = <T>(data:T) => requests('/user/edit',data)
export const deleteUser = <T>(data:T) => requests('/user/delete',data) 