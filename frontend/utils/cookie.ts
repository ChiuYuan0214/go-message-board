export const setCookies = (userId: number, token: string) => {
  const expirationTime = new Date(new Date().getTime() + 30 * 60 * 1000);
  const expiresGMT = expirationTime.toUTCString();

  document.cookie = `token=${token}; expires=${expiresGMT}; path=/`;
  document.cookie = `userId=${userId}; expires=${expiresGMT}; path=/`;
};

export const parseCookies = (cookies: string) => {
  const cookieMap: { [key: string]: string } = {};
  cookies.split(";").reduce((obj, cookie) => {
    const [key, val] = cookie.trim().split("=");
    obj[key] = val;
    return obj;
  }, cookieMap);
  return cookieMap;
};
