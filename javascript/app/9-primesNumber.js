const cekPrimes = (n) => {
  if (n == 1) return false;
  for (let i = 2; i < n; i++) {
    if (n % i == 0) return false;
  }
  return true;
};

const findPrimes = (n) => {
  let i = 1;
  let result = [];
  while (result.length < n) {
    if (cekPrimes(i)) result.push(i);
    i++;
  }
  return result.toString();
};

console.log("result: ", findPrimes(10));
