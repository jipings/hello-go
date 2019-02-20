 /**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
    let m = matrix.length;
    if(m ==0){
        return false
    }
    let n = matrix[0].length
    if (n== 0) {
        return false
    }
    let i = m - 1;
    let j = 0
 
    while (i>=0 && j < n) {
     if (matrix[i][j] == target) {
         return true
     } else if (matrix[i][j]<target) {
         j = j+1
     } else {
         i = i -1
     }
    }
    return false
 };