var io = require('../../helpers/js/io');
var input = io.readLines('../input/02.txt');

day2_1();
day2_2();

// A X Rock
// B Y Paper
// C Z Scissors

// A Rock 1
// B Paper 2
// C Scissors 3
// X lose 0
// Y draw 3
// Z win 6

function day2_1(){
    let total = 0;
    for(let i = 0; i < input.length; i++){
        switch(input[i][2]){
            case 'X':
                total += 1;
                if (input[i][0] == 'A'){
                    total += 3;
                    break;
                }
                else if (input[i][0] == 'B') {
                    break;
                }
                else {
                    total += 6;
                    break;
                }
            case 'Y':
                total += 2;
                if (input[i][0] == 'A'){
                    total += 6;
                    break;
                }
                else if (input[i][0] == 'B') {
                    total += 3;
                    break;
                }
                else {
                    break;
                }
            case 'Z':
                total += 3;
                if (input[i][0] == 'A'){
                    break;
                }
                else if (input[i][0] == 'B') {
                    total += 6;
                    break;
                }
                else {
                    total += 3;
                    break;
                }
        }
    }
    console.log("Puzzle 1:");
    console.log(total);
}

function day2_2(){
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
    console.log("Puzzle 2:");
    console.log(total);
}
