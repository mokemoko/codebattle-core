<script lang="ts">
  import { Table } from 'sveltestrap'
  import type { Match } from '../generated'

  export let matches: Match[]
</script>

<h3>History</h3>
<Table striped hover>
  <thead>
  <tr>
    <th>ID</th>
  </tr>
  </thead>
  <tbody>
  {#each matches as match}
    <tr>
      <td>{match.id}</td>
    </tr>
  {/each}
  </tbody>
</Table>

<style>
</style>
