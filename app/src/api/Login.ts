import requests from '@/server/request'
export const register = <T>(data:T) => requests('/Login/register',data)