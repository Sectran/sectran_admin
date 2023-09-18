import requests from '@/server/request'
export const register = <T>(data:T) => requests('/store/register',data)