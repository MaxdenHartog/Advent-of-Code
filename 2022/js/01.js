var io = require('./helpers/io');
var input = io.readLines('./input/01.txt');

day1_1();
day1_2();

function day1_1(){
    let mostCalories = 0;
    let totalCalories = 0;
        for(let i = 0; i < input.length; i++){
            if (input[i] != '') {
            totalCalories = totalCalories + Number(input[i]);            
            }
            else {
                if(mostCalories <= totalCalories){
                    mostCalories = totalCalories;
                }
                totalCalories = 0;
            }
        }

    console.log("Puzzle 1:");
    console.log(mostCalories);
}

function day1_2(){
    let mostCalories = 0;
    let totalCalories = 0;
    let elves =[];
        for(let i = 0; i < input.length; i++){
            if (input[i] != '') {
            totalCalories = totalCalories + Number(input[i]);            
            }
            else {
                elves.push(totalCalories);
                totalCalories = 0;
            }
        }
    
    elves.sort(function(a, b){return b-a});
    console.log("Puzzle 2:");
    console.log(elves[0] + elves[1] + elves[2]);
}