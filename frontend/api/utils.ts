const headers = new Headers();
headers.append("Content-Type", "application/json");

let token = "";
let tokenExpire = 0;

export const setLocalToken = (newToken: string) => {
  token = newToken;
  localStorage.setItem("token", newToken);
};

export const getLocalToken = () => token;
export const getLocalExpireTime = () => tokenExpire;

export const setLocalTokenExpire = (expireTime: number) => {
  tokenExpire = expireTime;
  localStorage.setItem("token-expire", expireTime + "");
};

export const setLocalUserId = (userId: number) => {
  localStorage.setItem("userId", userId + "");
};

export const getLocalUserId = () => {
  return +(localStorage.getItem("userId") || 0);
};

export const syncLocalStorage = () => {
  if (!token) {
    token = localStorage.getItem("token") || "";
  }
  if (!tokenExpire) {
    tokenExpire = +(localStorage.getItem("token-expire") || 0);
  }
};

export const isTokenValid = () => {
  return (
    tokenExpire &&
    new Date().getTime() - new Date(tokenExpire).getTime() > 5 * 60 * 1000
  );
};

export const removeLocalToken = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("token-expire");
  localStorage.removeItem("userId");
  token = "";
  tokenExpire = 0;
};

export const setBearerToken = (cusHeaders?: Headers) => {
  const h = cusHeaders || headers;
  h.set("Authorization", `Bearer ${token}`);
};

const api = {
  async get(url: string): Promise<any> {
    if (token) setBearerToken();
    const response = await fetch(url, {
      method: "GET",
      headers,
    });
    return this.resolve(response);
  },
  async post(
    url: string,
    body: { [key: string]: any },
    cusheaders?: Headers
  ): Promise<any> {
    if (token) setBearerToken();
    const response = await fetch(url, {
      method: "POST",
      headers,
      body: JSON.stringify(body),
    });
    return this.resolve(response);
  },
  async postForm(url: string, body: FormData): Promise<any> {
    const newHeaders = new Headers();
    if (token) setBearerToken(newHeaders);
    const response = await fetch(url, {
      method: "POST",
      headers: newHeaders,
      body,
    });
    return this.resolve(response);
  },
  async put(url: string, body: { [key: string]: any }): Promise<any> {
    if (token) setBearerToken();
    const response = await fetch(url, {
      method: "PUT",
      headers,
      body: JSON.stringify(body),
    });
    return this.resolve(response);
  },
  async delete(url: string, body: { [key: string]: any }): Promise<any> {
    if (token) setBearerToken();
    const response = await fetch(url, {
      method: "DELETE",
      headers,
      body: JSON.stringify(body),
    });
    return this.resolve(response);
  },
  async resolve(res: Response): Promise<any> {
    if (!res.ok) {
      console.log("status:", res.status);
    }
    return await res.json();
  },
};

export default api;
