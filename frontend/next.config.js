/** @type {import('next').NextConfig} */
const nextConfig = {
  basePath: "",
  reactStrictMode: true,
  swcMinify: true,
  publicRuntimeConfig: {
    GENERAL_IP: process.env.GENERAL_IP,
    SSR_GENERAL_IP: process.env.SSR_GENERAL_IP,
    SECURITY_IP: process.env.SECURITY_IP,
    SSR_SECURITY_IP: process.env.SSR_SECURITY_IP,
    IMAGE_PATH: process.env.IMAGE_PATH,
    SOCKET_URI: process.env.SOCKET_URI,
  },
};

module.exports = nextConfig;
