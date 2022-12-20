<script lang="ts">
  import { Button, Card, Col, Container, Modal, ModalBody, Row, Table } from 'sveltestrap'
  import dayjs from 'dayjs'
  import type { Match } from '../generated'
  import ScoreLabel from './atoms/ScoreLabel.svelte'
  import RankLabel from './atoms/RankLabel.svelte'
  import BombermanMatchView from '../bomberman/BombermanMatchView.svelte'
  import MatchModal from './MatchModal.svelte'
  import StatusBadge from './atoms/StatusBadge.svelte'
  import { userState } from "../domain/state";

  export let matches: Match[]

  let isOpenMatchEntry = false
  let selectedMatch: Match | null
</script>

<h3>
  <span>Matching History</span>
  {#if $userState}
    <Button outline on:click={() => isOpenMatchEntry = true}>Battle</Button>
  {/if}
</h3>

{#each matches as match}
  <div class="body">
    <Card>
      <Container>
        <Row>
          <Col xs="2">Time</Col>
          <Col>
            <div class="d-flex">
              <span class="flex-grow-1">{dayjs(match.createdAt).format('YYYY/MM/DD HH:mm:ss')}</span>
              <StatusBadge status={match.type}/>
            </div>
          </Col>
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
      {#if match.status === "finished"}
        <a class="stretched-link" href="#{match.id}" on:click={() => selectedMatch = match}></a>
      {/if}
    </Card>
  </div>
{/each}

<Modal size="lg" isOpen={selectedMatch} toggle={() => selectedMatch = null}>
  <ModalBody class="mx-auto">
    <BombermanMatchView matchId={selectedMatch.id} playerNames={selectedMatch.entries.map(e => e.name)}/>
  </ModalBody>
</Modal>

<MatchModal isOpen={isOpenMatchEntry} callback={() => isOpenMatchEntry = false}/>

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
