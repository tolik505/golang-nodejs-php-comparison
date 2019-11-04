var starttime = new Date().getTime();
console.log(starttime);

var array = [3,4,1,3,5,1,92,2,4124,424,52,12];

for (var c = 0; c < 1000000; c++) {
    for (var i = 0; i < array.length; i++) {
        for (var y = 0; y < array.length - 1; y++) {
            if (array[y+1] < array[y]) {
                var t = array[y];
                array[y] = array[y + 1];
                array[y + 1] = t;
            }
        }
    }
}
console.log(array);
console.log(new Date().getTime() - starttime + ' ms');

const used = process.memoryUsage().heapUsed / 1024;
console.log(`The script uses approximately ${Math.round(used * 100) / 100} KB`);