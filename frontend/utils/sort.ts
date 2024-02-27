export const mergeSortedList = (
  list1: any[],
  list2: any[],
  getter: (i: any) => number
) => {
  let index1 = 0;
  let index2 = 0;
  const result: any[] = [];
  while (index1 < list1.length && index2 < list2.length) {
    if (getter(list1[index1]) < getter(list2[index2])) {
      result.push(list1[index1]);
      index1++;
    } else {
      result.push(list2[index2]);
      index2++;
    }
  }

  while (index1 < list1.length) {
    result.push(list1[index1++]);
  }

  while (index2 < list2.length) {
    result.push(list2[index2++]);
  }
  return result;
};
