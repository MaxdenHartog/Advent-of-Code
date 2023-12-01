var io = require('../../helpers/js/io');
var input = io.readLines('../input/test.txt');

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

day1_1();
day1_2();

function day1_1(){
    let sum = 0;
    for (let i = 0; i < input.length; i++){
        let calibration = input[i].match(/\d/g)
        calibration = calibration[0] + calibration[calibration.length - 1]
        sum += parseInt(calibration);
    }
    console.log("Puzzle 1:");
    console.log(sum);
}

function day1_2(){
    let sum = 0;
    for (let i = 0; i < input.length; i++){
        let calibration = input[i].match(/\d|(zero)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)/g);
        let first = !isNaN(parseInt(calibration[0])) ?
            parseInt(calibration[0]) * 10 :
            numbers.indexOf(calibration[0]) * 10;
        let last = !isNaN(parseInt(calibration[calibration.length - 1])) ?
            parseInt(calibration[calibration.length - 1]) :
            numbers.indexOf(calibration[calibration.length - 1]);
        sum += first + last;
    }
    console.log("Puzzle 2:");
    console.log(sum);
}