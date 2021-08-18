const arrayPoints = [
  [34.79600143432617, 700.9229736328125],
  [71.913002014160167, 737.7329711914062],
  [426.864990234375, 737.7670288085938],
  [425.6109924316406, 529.6099853515625],
  [387.1189880371094, 528.7100219726562],
  [386.4949951171875, 448.89599609375],
];

const convertToStrong = (index, arrayPoint, stringPoint = "") => {
  let = newStringPoint = stringPoint;
  if (index === 0) {
    return newStringPoint;
  } else {
    arrayPoint[arrayPoint.length - index].map((item) => {
      newStringPoint = newStringPoint.concat(item);
      if (index > 1) newStringPoint = newStringPoint.concat(", ");
    });
  }

  return convertToStrong(index - 1, arrayPoint, newStringPoint);
};

const result = convertToStrong(arrayPoints.length, arrayPoints);

console.log(result);
