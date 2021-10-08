<?php
//hacktoberfest

$data = [
    'name' => 'Yoppy Dimas',
    'region' => 'ID'
];

//encode array return string
$encode = json_encode($data);
echo $encode;

//decode json return array
$array = json_decode($encode, true, JSON_PRETTY_PRINT);
print_r($array);