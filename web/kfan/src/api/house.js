import request from './index'

export function getHouses(){
    return request({
        "method": "get",
        "url": "house/v1/houses"
    })
}