<?php

  function factorial(int $num) {
    if($num === 0) {
      return 1;
    }
    return $num * factorial($num - 1);
  }

  $result = factorial(5);
  echo "result: " . $result . "\n";
?>