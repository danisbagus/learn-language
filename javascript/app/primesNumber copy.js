formatTimeElapsed = (time) => {
  if (time && time >= 0) {
    const days = (time - (time % 1440)) / 1440;
    const hours = ((time % 1440) - (time % 60)) / 60;
    const mins = time % 60;
    return `${days ? days + " days " : ""}${hours ? hours + " hours " : ""}
    ${mins ? mins + " mins " : ""}`;
  } else {
    return "-";
  }
};

console.log("time :", formatTimeElapsed(120));
// console.log("time ", 60 * 24);
