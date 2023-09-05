import { paths } from "@/types/openapi"
import createClient from "openapi-fetch"
import { v4 } from "uuid"

// TODO(SP) fetch from env
const baseUrl = "http://localhost:3333"

const uuid = (): string => v4()

async function customFetch(
  input: RequestInfo | URL,
  init?: RequestInit
): Promise<Response> {
  const result = await fetch(input, init)
  if (result.status !== 200) {
    console.log(init?.method, input, `unexpected status ${result.status}`)
    const responseBody = await result.clone()
    const responseBodyJson = await responseBody.json()
    console.log(responseBodyJson)
  }
  if (result.status === 401) {
    // TODO(SP) redirect to login
  }
  return result
}

export const { GET, POST, PUT, DELETE, PATCH } = createClient<paths>({
  baseUrl: baseUrl,
  cache: "no-cache",
  headers: {
    "Content-Type": "application/json",
    "sherpa-session-id": "dummy-session-id",
    "sherpa-request-id": uuid(),
  },
  fetch: customFetch,
})
