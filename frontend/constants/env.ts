import getConfig from "next/config";

const { publicRuntimeConfig: config } = getConfig();

export const GENERAL_IP = (() =>
  config && config.GENERAL_IP ? config.GENERAL_IP : "http://localhost:4444")();
export const SSR_GENERAL_IP = (() =>
  config && config.SSR_GENERAL_IP
    ? config.SSR_GENERAL_IP
    : "http://localhost:4444")();
export const SECURITY_IP = (() =>
  config && config.SECURITY_IP
    ? config.SECURITY_IP
    : "http://localhost:5555")();
export const SSR_SECURITY_IP = (() =>
  config && config.SSR_SECURITY_IP
    ? config.SSR_SECURITY_IP
    : "http://localhost:5555")();
export const IMAGE_PATH = (() =>
  config && config.IMAGE_PATH
    ? config.IMAGE_PATH
    : "http://127.0.0.1:5555/uploads/images/")();
export const SOCKET_URI = (() =>
  config && config.SOCKET_URI
    ? config.SOCKET_URI
    : "ws://127.0.0.1:6666/chat")();
