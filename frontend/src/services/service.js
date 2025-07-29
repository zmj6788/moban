import axios from "axios";
import {useStore} from "@/stores/store";

export const Service = axios.create({
    timeout: 7000,
    baseURL: "",
    headers: {
        "Content-Type": "application/json"
    }
})


Service.interceptors.request.use(request => {
    // 一般用于用户的token
    const store = useStore()
    request.headers["token"] = store.userInfo.token
    return request
})


Service.interceptors.response.use(response => {
    return response.data
})