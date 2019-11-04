<?php
$starttime = microtime(true);

$array = [3,4,1,3,5,1,92,2,4124,424,52,12];

for ($c = 0; $c < 1000000; $c++) {
    for ($i = 0; $i < count($array); $i++) {
        for ($y = 0; $y < count($array)-1; $y++) {
            if ($array[$y+1] < $array[$y]) {
                $t = $array[$y];
                $array[$y] = $array[$y+1];
                $array[$y+1] = $t;
            }
        }
    }
}
print_r($array);
echo (microtime(true) - $starttime) * 1000 . ' ms' . PHP_EOL;
print_mem();

function print_mem()
{
   /* Currently used memory */
   $mem_usage = memory_get_usage();
   
   /* Peak memory usage */
   $mem_peak = memory_get_peak_usage();

   echo 'The script is now using: ' . round($mem_usage / 1024) . 'KB of memory.' . PHP_EOL;
   echo 'Peak usage: ' . round($mem_peak / 1024) . 'KB of memory.';
}