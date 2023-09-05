import { GET } from "@/lib/api"

async function listSprints(projectId: number) {
  const { data, response } = await GET("/sprints", {
    params: {
      query: {
        project_id: projectId,
      },
    },
  })

  if (response.status !== 200) {
    throw new Error(
      `Failed to fetch sprints, unexpected status ${response.status}`
    )
  }

  return data
}

export default async function Home() {
  const data = await listSprints(1)

  return (
    <main className="">
      {data?.map((sprint) => <p key={sprint.id}>{sprint.title}</p>)}
    </main>
  )
}
