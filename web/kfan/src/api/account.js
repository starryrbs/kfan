import request from './index'

export function Login(username){
    return request({
        "method": "post",
        "url": "account/v1/login",
        data: {username}
    })
}