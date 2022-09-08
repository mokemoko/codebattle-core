<script lang="ts">
  import { Container, Row, Col } from 'sveltestrap'
  import { ContestsApi, Configuration, Contest } from './generated'
  import LeaderBoard from './components/LeaderBoard.svelte'
  import Status from './components/Status.svelte'
  import Pickup from './components/Pickup.svelte'
  import History from './components/History.svelte'

  async function fetchData(): Promise<Contest> {
    const client = new ContestsApi(new Configuration({
      basePath: 'http://localhost:8081',
    }))
    const contests = await client.getContests()
    return client.getContestById({ contestId: contests[0].id })
  }

  const promise = fetchData()
</script>

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
        <Pickup/>
        <LeaderBoard entries={contest.ranking}/>
      </Col>
      <Col xs="4">
        <Status/>
        <History matches={contest.recentMatches}/>
      </Col>
    </Row>
  </Container>
{/await}

<style>
</style>
