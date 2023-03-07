
import Layout from '@/components/Layout.js'
import '@/styles/globals.css'
import { ChakraProvider } from '@chakra-ui/react'


export default function App({ Component, pageProps }) {
  return (
    <ChakraProvider>
    <Layout>
      <Component {...pageProps} />
    </Layout>
    </ChakraProvider>

  )
}
