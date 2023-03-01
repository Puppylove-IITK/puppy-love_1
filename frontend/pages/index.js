import Head from 'next/head'
import { useRouter } from 'next/router'
import { Inter } from 'next/font/google'
import Login from '@/components/login'
import styles from '@/styles/Home.module.css'

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <>
      <Head>
        <title>Puppylove | Login</title>
      </Head>
      <div className={styles.container}>
        <Login />
      </div>

    </>
  )
}
