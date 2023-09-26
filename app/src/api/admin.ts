import requests from '@/server/request'
export const addDepartment = <T>(data:T) => requests('/department/add',data)