const {readFileSync} = require('fs');

module.exports = {
    readLines: function(file){
        const contents = readFileSync(file, 'utf-8');
        const array = contents.split(/\r?\n/);
        return array
    }
};