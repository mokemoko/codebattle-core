<script lang="ts">
  import { Button, Card, Col, Container, Row, Table } from 'sveltestrap'
  import dayjs from 'dayjs'
  import type { Match } from '../../generated'

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
          <Col xs="3">Time</Col>
          <Col>{dayjs(match.createdAt).format()}</Col>
        </Row>
        <Row>
          <Col xs="3">Result</Col>
          <Col>
            {#each match.entries as entry, index}
              <div>{entry.rank} {entry.name} {entry.afterScore} {entry.afterScore - entry.beforeScore}</div>
            {/each}
          </Col>
        </Row>
      </Container>
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
</style>
