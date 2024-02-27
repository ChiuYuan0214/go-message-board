import Layout from "@/layout";
import "../styles/globals.css";
import type { AppProps } from "next/app";
import { ChatProvider } from "@/context/chat";
import { AuthProvider } from "@/context/auth";

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <AuthProvider>
      <ChatProvider>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </ChatProvider>
    </AuthProvider>
  );
}

export default MyApp;
