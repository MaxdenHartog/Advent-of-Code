var fs = require('fs');
var array = fs.readFileSync('input.txt').toString().split("\n");

// A Rock 1
// B Paper 2
// C Scissors 3
// X lose 0
// Y draw 3
// Z win 6
let total = 0;
for(i in array) {
    switch(array[i][2]){
        case 'X':
            if(array[i][0] == 'A'){
                total += 3;
                break;
            }
            else if(array[i][0] == 'B'){
                total += 1;
                break;
            }
            else {
                total += 2;
                break;
            }
        case 'Y':
            total += 3;
            if(array[i][0] == 'A'){
                total += 1;
                break;
            }
            else if(array[i][0] == 'B'){
                total += 2;
                break;
            }
            else {
                total += 3;
                break;
            }
        case 'Z':
            total += 6;
            if(array[i][0] == 'A'){
                total += 2;
                break;
            }
            else if(array[i][0] == 'B'){
                total += 3;
                break;
            }
            else {
                total += 1;
                break;
            }        
    }
}
console.log(total);