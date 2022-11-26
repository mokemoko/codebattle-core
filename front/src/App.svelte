<script lang="ts">
  import { Container, Row, Col } from 'sveltestrap'
  import { ContestsApi, Configuration, Contest, UsersApi } from './generated'
  import LeaderBoard from './components/LeaderBoard.svelte'
  import Status from './components/Status.svelte'
  import Pickup from './components/Pickup.svelte'
  import History from './components/History.svelte'
  import Navigation from './components/Navigation.svelte'
  import { userState } from './state/user'

  async function fetchData(): Promise<Contest> {
    const userClient = new UsersApi(new Configuration({
      basePath: 'http://localhost:8081',
      credentials: 'include',
    }))
    try {
      const user = await userClient.getMe()
      userState.set(user)
    } catch (e) {
      console.debug(e)
    }

    const client = new ContestsApi(new Configuration({
      basePath: 'http://localhost:8081',
    }))
    const contests = await client.getContests()
    return client.getContestById({ contestId: contests[0].id })
  }

  const promise = fetchData()
</script>

<Navigation/>
{#await promise}
  <div class="d-flex justify-content-center align-items-center" style="height: 100%">
    <div class="spinner-border" role="status"></div>
  </div>
{:then contest}
  <Container>
    <Row>
      <h1>{contest.name}</h1>
    </Row>
    <Row>
      <Col xs="8">
        <Pickup match={contest.recentMatches[0]}/>
        <LeaderBoard entries={contest.ranking}/>
      </Col>
      <Col xs="4">
        <Status user={null}/>
        <History matches={contest.recentMatches}/>
      </Col>
    </Row>
  </Container>
{/await}

<style>
</style>
