// 给你一个整数数组 arr，如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。
function uniqueOccurrences(arr: number[]): boolean {
    let map = new Map<number, number>()
    for (let i = 0; i < arr.length; i++) {
        map.set(arr[i], (map.get(arr[i]) ?? 0)+1)
    }
    let check = new Map<number, boolean>()
    for (let [_, value] of map){
        if (check.get(value)){
            return false
        }else {
            check.set(value,true)
        }
    }
    return true
}

const test: number[] = [1,2,2,1,1,3]
console.log(uniqueOccurrences(test))