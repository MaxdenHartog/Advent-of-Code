var io = require('../../helpers/js/io');
var input = io.readLines('../input/01.txt');

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

const numbersReverse = [
    "orez",
    "eno",
    "owt",
    "eerht",
    "ruof",
    "evif",
    "xis",
    "neves",
    "thgie",
    "enin"
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
        let calibrationReverse = input[i].split("").reverse().join("").match(/\d|(orez)|(eno)|(owt)|(eerht)|(ruof)|(evif)|(xis)|(neves)|(thgie)|(enin)/g);
        let last = !isNaN(parseInt(calibrationReverse[0])) ?
            parseInt(calibrationReverse[0]) :
            numbersReverse.indexOf(calibrationReverse[0]);
        sum += first + last;
    }
    console.log("Puzzle 2:");
    console.log(sum);
}