const s = "asdasdasd12";

function ss (strings) {
    const arr = strings.split('')
    let str = '';
    let strArr = [];
    let strMap= {};

    arr.forEach((t, i) => {
        if(strMap[t] !== undefined) {
            strArr.push(str)
            str = t
            strMap = {}
            strMap[t]=i  
        } else {
            str += t 
            strMap[t]=i
            if(i == arr.length - 1) {
                strArr.push(str)
            }
        }
    })
    return strArr
}

console.log(ss(s))