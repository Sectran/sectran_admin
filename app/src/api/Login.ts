import requests from '@/server/request'
export const login = <T>(data: T) => requests('/login/index', data)

export const userInfo = () => requests('/userInfo', {}, 'get')