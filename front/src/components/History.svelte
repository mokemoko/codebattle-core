<script lang="ts">
  import { Button, Card, Col, Container, Row, Table } from 'sveltestrap'
  import dayjs from 'dayjs'
  import type { Match } from '../generated'
  import ScoreLabel from './atoms/ScoreLabel.svelte'
  import RankLabel from './atoms/RankLabel.svelte'

  export let matches: Match[]
</script>

<h3>
  <span>Matching History</span>
  <Button outline>Battle</Button>
</h3>

{#each matches as match}
  <div class="body">
    <Card>
      <Container>
        <Row>
          <Col xs="2">Time</Col>
          <Col>{dayjs(match.createdAt).format('YYYY/MM/DD HH:mm:ss')}</Col>
        </Row>
        <Row>
          <Col xs="2">Result</Col>
          <Col>
            {#each match.entries as entry, index}
              <div class="d-flex">
                <RankLabel rank={entry.rank}/>
                <span class="flex-grow-1 mx-2">{entry.name}</span>
                <ScoreLabel score={entry.afterScore} delta={entry.afterScore - entry.beforeScore}/>
              </div>
            {/each}
          </Col>
        </Row>
      </Container>
      <a href="#{match.id}" class="stretched-link"></a>
    </Card>
  </div>
{/each}

<style>
  h3 {
    display: flex;
    justify-content: space-between;
  }

  .body {
    margin-bottom: 8px;
  }

  .body :global(.container) {
    padding-right: 0.5em;
    padding-left: 0.5em;
  }
</style>
