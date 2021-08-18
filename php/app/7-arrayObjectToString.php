<?php

$arrayPoint = array (
  array("x"=> 34.79600143432617, "y"=> 700.9229736328125),
  array("x"=> 71.913002014160167, "y"=> 737.7329711914062),
  array("x"=> 426.864990234375, "y"=> 737.7670288085938),
  array("x"=> 425.6109924316406, "y"=> 529.6099853515625),
  array("x"=> 387.1189880371094, "y"=> 528.7100219726562),
  array("x"=> 386.4949951171875, "y"=> 448.89599609375),
);

  $arrayLength = sizeOf($arrayPoint);
  $stringPoint = '';

  for($i=0; $i < $arrayLength; $i++) {
    $stringPoint .= $arrayPoint[$i]["x"] . ", " . $arrayPoint[$i]["y"];
    if($i < $arrayLength - 1) $stringPoint .= ", ";
  }

  print_r($stringPoint);

?>