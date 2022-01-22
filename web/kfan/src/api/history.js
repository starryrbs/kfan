import request from "@/api/index";

export function saveHistory(userId, objId, objType) {
    return request({
        "method": "post",
        "url": "history/v1/history",
        data: {obj_id: objId, obj_type: objType, user_id: userId}
    })
}