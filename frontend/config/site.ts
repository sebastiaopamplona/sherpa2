import { NavItem } from "@/types/nav"

interface SiteConfig {
  name: string
  description: string
  landingPageNav: NavItem[]
  applicationNav: NavItem[]
  links: {
    twitter: string
    github: string
    docs: string
  }
  registration: {
    enabled: boolean
    link?: string
    fallback: string
    fallbackText: string
  }
}

export const siteConfig: SiteConfig = {
  name: "Sherpa",
  description: "Visualize your work",
  landingPageNav: [
    {
      title: "About",
      href: "#about",
    },
  ],
  applicationNav: [
    {
      title: "Time keeper",
      href: "#time-keeper",
    },
    {
      title: "Sprint",
      href: "#sprint",
    },
  ],
  links: {
    twitter: "https://twitter.com/shadcn",
    github: "https://github.com/shadcn/ui",
    docs: "https://ui.shadcn.com",
  },
  registration: {
    enabled: true,
    fallback: "https://google.com",
    fallbackText: "In progress",
  },
}
