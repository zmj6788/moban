import {Service} from "@/services/service";


// 邮箱登录接口
export function emailLoginApi(data) {
    return Service.post("/api/email_login", data)
}

// 注销
export function logoutApi() {
    return Service.post("/api/logout")
}