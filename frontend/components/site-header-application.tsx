import { MainNav } from "@/components/main-nav"
import { siteConfig } from "@/config/site"
import React from "react"

export function SiteHeaderApplication() {
  return (
    <header className="shadow-level1 dark:shadow-dark-level1 sticky top-0 z-40 w-full border-b border-b-slate-200 bg-white dark:border-b-slate-700 dark:bg-slate-900 ">
      <div className="container flex h-16 items-center space-x-4 sm:justify-between sm:space-x-0">
        <MainNav items={siteConfig.applicationNav} />
        <div className="flex flex-1 items-center justify-end space-x-4">
          <nav className="flex items-center space-x-1"></nav>
        </div>
      </div>
    </header>
  )
}
