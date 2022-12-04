var io = require('../../helpers/js/io');
var input = io.readLines('../input/04.txt');

day4_1();
day4_2();

function day4_1(){
    let total = 0;
    for(let i = 0; i < input.length; i++){
        var bounds = input[i].split(/(?:,|-)+/).map(Number);
        if ((bounds[0] <= bounds[2] && bounds[1] >= bounds[3]) || (bounds[2] <= bounds[0] && bounds[3] >= bounds[1])) {
            total++;
        }
    }
    console.log("Puzzle 1:");
    console.log(total);
}

function day4_2(){
    let total = 0;
    for(let i = 0; i < input.length; i++){
        var bounds = input[i].split(/(?:,|-)+/).map(Number);
        if ((bounds[0] <= bounds[2] && bounds[1] >= bounds[2]) || (bounds[2] <= bounds[0] && bounds[3] >= bounds[0])) {
            total++;
        }
    }
    console.log("Puzzle 2:");
    console.log(total);
}