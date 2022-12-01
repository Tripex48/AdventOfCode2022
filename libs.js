const formatRuntime = (ms) => {
  // You can customize this function to your liking

  // microseconds
  if (ms <= 1) {
    return `${Math.round(ms * 1000)}µs`;
  }
  // miliseconds
  if (ms < 1000) {
    const wholeMs = Math.floor(ms);
    return `${wholeMs}ms ${formatRuntime(ms - wholeMs)}`;
  }
  const sec = ms / 1000;
  if (sec < 60) {
    const wholeSec = Math.floor(sec);
    const remMs = ms - wholeSec * 1000;
    return `${wholeSec}s ${formatRuntime(remMs)}`;
  }
  // Minutes (hopefully it doesn't get to this point lol)
  return `${Math.floor(sec / 60)}m ${formatRuntime((sec % 60) * 1000)}`;
};

module.exports = {
  formatRuntime,
};
