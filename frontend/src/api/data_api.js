import {Service} from "@/services/service";

export function getWeekDataApi(){
    /*
{
  "code": 0,
  "data": {
    "date_list": ["2023-03-12","2023-03-13","2023-03-14","2023-03-15","2023-03-16","2023-03-17","2023-03-18"],
    "login_data": [1,0,1,2,1,1,1],
    "sign_data": [0,0,0,0,0,0,0]
  },
  "msg": "成功"
}
     */
    return Service.get("/api/data_login")
}

export function getSumDataApi(){
    /*
{
  "code": 0,
  "data": {
    "user_count": 6,
    "article_count": 9,
    "message_count": 11,
    "chat_group_count": 14,
    "now_login_count": 1,
    "now_sign_count": 0
  },
  "msg": "成功"
}
*/
    return Service.get("/api/data_sum")
}