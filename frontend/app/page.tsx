import { GET } from "@/lib/api"

async function listSprints() {
  const { data, response } = await GET("/sprints", {
    params: {
      query: {
        project_id: 1,
      }
    }
  })

  return data
}

export default async function Home() {
  const data = await listSprints()

  return (
    <main className="">
      {data?.map((sprint) => <p key={sprint.id}>{sprint.title}</p>)}
    </main>
  )
}
