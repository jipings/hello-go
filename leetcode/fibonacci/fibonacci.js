
function fibonacci(n) {
    if(n<2) {
        return n
    }
    return fibonacci(n-2) + fibonacci(n-1)
}

console.time("start")

for(let i= 0;i<30;i++) {
    console.log(fibonacci(i))
}

console.timeEnd("start")