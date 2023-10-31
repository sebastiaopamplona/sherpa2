import "./globals.css"
import { SiteHeaderApplication } from "@/components/site-header-application"
import Head from "next/head"

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <SiteHeaderApplication />
      <main>
        <Head>
          <title>Next.js</title>
          <meta
            name="description"
            content="Next.js template for building apps with Radix UI and Tailwind CSS"
          />
          <meta name="viewport" content="width=device-width, initial-scale=1" />
          <link rel="icon" href="/favicon.ico" />
        </Head>
        <section className="container grid items-center pt-6 pb-8">
          {children}
        </section>
      </main>
    </>
  )
}
