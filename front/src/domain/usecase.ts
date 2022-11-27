import { contestClient, entryClient, userClient } from './api'
import { contestState, userState } from './state'
import { get } from 'svelte/store'

export async function loadData() {
  try {
    const user = await userClient.getMe()
    userState.set(user)
  } catch (e) {
    console.debug(e)
  }

  const contests = await contestClient.getContests()
  const contest = await contestClient.getContestById({ contestId: contests[0].id })
  contestState.set(contest)
}

export async function registerEntry(name: string, repository: string) {
  const contestId = get(contestState).id
  await entryClient.postEntry({ postEntryRequest: { contestId, name, repository } })

  const updated = await contestClient.getContestById({ contestId })
  contestState.set(updated)
}
