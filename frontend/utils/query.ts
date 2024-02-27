export const getParam = (
  param: string | string[] | undefined,
  fallback: string
) => {
  return Array.isArray(param) || !param ? fallback : param;
};
