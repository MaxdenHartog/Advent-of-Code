var io = require('../../helpers/js/io');
var input = io.readLines('../input/03.txt');

day3_1();
day3_2();

// a: .charCodeAt(0) - 96
// A: .charCodeAt(0) - 38

function day3_1(){
    let totalPriority = 0;
        for(let i = 0; i < input.length; i++){
            var compartments = [input[i].slice(0, input[i].length / 2), input[i].slice(input[i].length/ 2)];
            for(let c in compartments[0]){
                if (compartments[1].includes(compartments[0][c])) {
                    totalPriority += ((compartments[0][c].charCodeAt(0) - 96 < 0) ? compartments[0][c].charCodeAt(0) - 38 : compartments[0][c].charCodeAt(0) - 96);
                    break;
                }
            }
        }
    console.log('Puzzle 1:');
    console.log(totalPriority);
}

function day3_2(){
    let totalPriority = 0;
        for(let i = 0; i < input.length - 2; i+=3){
            for(let c in input[i]){
                if (input[i+1].includes(input[i][c]) && input[i+2].includes(input[i][c])) {
                    totalPriority += ((input[i][c].charCodeAt(0) - 96 < 0) ? input[i][c].charCodeAt(0) - 38 : input[i][c].charCodeAt(0) - 96);
                    break;
                }
            }
        }
    console.log('Puzzle 2:');
    console.log(totalPriority);
}
