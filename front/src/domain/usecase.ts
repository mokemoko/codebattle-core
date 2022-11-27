import { contestClient, userClient } from './api'
import { contestState, userState } from './state'

export async function loadData () {
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
