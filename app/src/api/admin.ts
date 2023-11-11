import requests from `@/server/request`
let prefix = "/sectran"
//部门
export const listDepartment =<T>(data:T)=> requests(`${prefix}/dept/list`,data)
export const addDepartment = <T>(data:T) => requests(`${prefix}/dept/add`,data)
export const editDepartment = <T>(data:T) => requests(`${prefix}/dept/edit`,data)
export const deleteDepartment = <T>(data:T) => requests(`${prefix}/dept/delete`,data) 

//用户
export const listUser =<T>(data:T)=> requests(`${prefix}/user/list`,data,`get`)
export const adduser = <T>(data:T) => requests(`${prefix}/user/add`,data)
export const edituser = <T>(data:T) => requests(`${prefix}/user/edit`,data)
export const deleteUser = <T>(data:T) => requests(`${prefix}/user/delete`,data) 

//角色
export const listRole =<T>(data:T)=> requests(`${prefix}/role/list`,data)
export const addRole =<T>(data:T)=> requests(`${prefix}/role/add`,data)
export const editRole =<T>(data:T)=> requests(`${prefix}/role/edit`,data)
export const deleteRole =<T>(data:T)=> requests(`${prefix}/role/delete`,data)






