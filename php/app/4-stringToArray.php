<?php
  $stringPoint =  "34.79600143432617 700.9229736328125 71.91300201416016 737.7329711914062 426.864990234375 737.7670288085938 425.6109924316406 529.6099853515625 387.1189880371094 528.7100219726562 386.4949951171875 448.89599609375 418.27301025390625 379.2030029296875 439.260009765625 388.7919921875 457.02398681640625 351.53900146484375 437.07598876953125 341.7040100097656 526.0599975585938 148.0590057373047 479.1059875488281 126.45800018310547 501.2149963378906 79.98300170898438 355.6340026855469 11.78499984741211";

  $arrayPoint = explode(" ", $stringPoint);

  $arrayLength = sizeof($arrayPoint);

  $coordinates = [];
  $points = [];

  for($i=0; $i < $arrayLength; $i++) {
    array_push($points, $arrayPoint[$i]);
    if($i % 2 !== 0) {
      array_push($coordinates, $points);
      $points = [];
    }
  }
  print_r($coordinates);
  print_r($coordinates[0]);

?>