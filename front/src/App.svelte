<script lang="ts">
  import { Container, Row, Col } from 'sveltestrap'
  import LeaderBoard from './components/LeaderBoard.svelte'
  import Status from './components/Status.svelte'
  import Pickup from './components/Pickup.svelte'
  import History from './components/History.svelte'
  import Navigation from './components/Navigation.svelte'
  import { loadData } from './domain/usecase'
  import { contestState } from './domain/state'
  import { onMount } from 'svelte'

  onMount(() => loadData())
</script>

<Navigation/>
{#if $contestState}
  <Container>
    <Row>
      <h1>{$contestState.name}</h1>
    </Row>
    <Row>
      <Col xs="8">
        <Pickup match={$contestState.recentMatches[0]}/>
        <LeaderBoard entries={$contestState.ranking}/>
      </Col>
      <Col xs="4">
        <Status user={null}/>
        <History matches={$contestState.recentMatches}/>
      </Col>
    </Row>
  </Container>
{:else}
  <div class="d-flex justify-content-center align-items-center" style="height: 100%">
    <div class="spinner-border" role="status"></div>
  </div>
{/if}

<style>
</style>
