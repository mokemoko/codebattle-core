<script lang="ts">
  import { Container, Row, Col } from 'sveltestrap'
  import { ContestsApi, Configuration, PublicEntry, Match } from './generated'
  import LeaderBoard from './components/LeaderBoard.svelte'
  import Status from "./components/Status.svelte";
  import Pickup from './components/Pickup.svelte'
  import History from "./components/History.svelte";

  let entries: PublicEntry[] = []
  let matches: Match[] = []

  async function fetchData() {
    const client = new ContestsApi(new Configuration({
      basePath: 'http://localhost:8081',
    }))
    const contests = await client.getContests()
    const contest = await client.getContestById({ contestId: contests[0].id })
    entries = contest.ranking
    matches = contest.recentMatches
  }

  const promise = fetchData()
</script>

<Container>
  <Row>
    <h1>CodeBattle</h1>
  </Row>
  <Row>
    <Col xs="8">
      <Pickup/>
      <LeaderBoard entries={entries}/>
    </Col>
    <Col xs="4">
      <Status />
      <History matches={matches} />
    </Col>
  </Row>
</Container>

<style>
</style>
