import {Service} from "@/services/service";


export function getArticleCalendarApi(){
    return Service.get("/api/articles/calendar")
}
