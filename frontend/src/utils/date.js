export function getFormatDateTime(dateStr) {
    let date = new Date(dateStr);
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let day = date.getDate();
    let hour = date.getHours();
    let minutes = date.getMinutes();
    let seconds = date.getSeconds();
    month = (month<10)? '0'+ month : month;
    day = (day<10)? '0'+ day : day;
    hour = (hour<10)? '0'+ hour : hour;
    minutes = (minutes<10)? '0'+ minutes : minutes;
    seconds = (seconds<10)? '0'+ seconds : seconds;
    let currentDate = year + "-" + month + "-" + day
            + " " + hour + ":" + minutes + ":" + seconds;
    return currentDate;
}

export function getFormatDate(dateStr, joins) {
    let date = new Date(dateStr);
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let day = date.getDate();
    month = (month<10)? '0'+ month : month;
    day = (day<10)? '0'+ day : day;

    if (joins === undefined){
        return year + "-" + month + "-" + day;
    }
    return year + "年" + month + "月" + day + "日";


}

