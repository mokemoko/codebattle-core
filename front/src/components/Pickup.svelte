<script lang="ts">
  import MatchView from './Match.svelte'
  import type { Match } from '../../generated'

  export let match: Match
</script>

<h3>Pickup</h3>
<MatchView match={match}/>

<style>
</style>
