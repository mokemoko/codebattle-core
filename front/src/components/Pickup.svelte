<script lang="ts">
  import BombermanMatchView from '../bomberman/BombermanMatchView.svelte'
  import type { Match } from '../generated'
  import { get } from 'svelte/store'
  import { contestState } from '../domain/state'

  $: match = get(contestState).recentMatches.find(match => match.type === "rated" && match.status === "finished")
</script>

{#if match}
  <h3>Pickup</h3>
  <!-- TODO: 汎用化 -->
  <BombermanMatchView matchId={match.id} playerNames={match.entries.map(e => e.name)}/>
{/if}

<style>
</style>
