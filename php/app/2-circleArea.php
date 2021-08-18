<<?php
$diameter = 20;
$radius = 10.0;
define("PHI", 3.14);
$isUsingRadius = true;
$error = null;

function calculate1(int $diameter) {
    $radius = $diameter / 2;
    return $radius * $radius * 3.14;
  }

function calculate2(int $radius) {
    return $radius * $radius * 3.14;
}

if($isUsingRadius)  {
$result = calculate2($radius);
} else {
$result = calculate1($diameter);
}

echo $result

?>