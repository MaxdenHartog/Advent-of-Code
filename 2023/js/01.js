var io = require('../../helpers/js/io');
var input = io.readLines('../input/01.txt');

day1_1();
day1_2();

function day1_1(){
    let sum = 0;
    for (let i = 0; i < input.length; i++){
        let calibtration = input[i].match(/\d/g)
        calibtration = calibtration[0] + calibtration[calibtration.length - 1]
        sum += parseInt(calibtration);
    }
    console.log("Puzzle 1:");
    console.log(sum);
}

function day1_2(){
    let sum = 0;
    for (let i = 0; i < input.length; i++){
        numbers.some(function(n) {return input[i].includes(n)})
    }
    console.log("Puzzle 2:");
    console.log(sum);
}

const numbers = [
    "zero",
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine"
]