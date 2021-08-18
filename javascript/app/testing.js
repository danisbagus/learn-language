const cekPrima = (n) => {
  var max = Math.sqrt(n);
  for (var i = 2; i <= max; i++) {
    if (n % i == 0) return false;
  }
  return true;
};

const getPrima = (cnt) => {
  const result = [];

  for (let n = 1; n < 10; n++) {
    if (cekPrima(n)) {
      result.push(n);
    }

    if (result.length >= cnt) {
      break;
    }
  }
  return result;
};

const convertArray = (input) => {
  let stringData = "";
  for (let i = 0; i < input.length; i++) {
    stringData = stringData.concat(input[i]);
    if (i < input.length - 1) stringData = stringData.concat(",");
  }

  return stringData;
};

const getBilanganPrima = (n) => {
  const arrayResult = getPrima(n);
  return convertArray(arrayResult);
};

console.log(getBilanganPrima(4));
