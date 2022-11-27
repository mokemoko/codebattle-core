<script lang="ts">
  import BombermanMatchView from '../bomberman/BombermanMatchView.svelte'
  import type { Match } from '../generated'

  export let match: Match
</script>

<h3>Pickup</h3>
<!-- TODO: 汎用化 -->
<BombermanMatchView matchId={match.id}/>

<style>
</style>
