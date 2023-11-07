import requests from '@/server/request'
export const login = <T>(data: T) => requests('/sectran/auth/login', data)