export const formatDate = (date: string) => {
  const d = new Date(date);

  return (
    padWithChar(d.getDate().toString(), 2, '0') +
    '.' +
    padWithChar((d.getMonth() + 1).toString(), 2, '0') +
    '.' +
    d.getFullYear() +
    ' - ' +
    padWithChar(d.getHours().toString(), 2, '0') +
    ':' +
    padWithChar(d.getMinutes().toString(), 2, '0')
  );
};

export const padWithChar = (s: string, l: number, c: string) => {
  while (s.length < l) {
    s = c + s;
  }
  return s;
};
