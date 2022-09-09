<script lang="ts">
  import { Table } from 'sveltestrap'
  import type { PublicEntry } from '../../generated'

  export let entries: PublicEntry[]
</script>

<h3>Leaderboard</h3>
<Table striped hover>
  <thead>
  <tr>
    <th>#</th>
    <th>EntryName</th>
    <th>UserName</th>
    <th>Score</th>
  </tr>
  </thead>
  <tbody>
  {#each entries as entry, i}
    <tr>
      <th scope="row">{i + 1}</th>
      <td>{entry.name}</td>
      <td>{entry.user.name}</td>
      <td>{entry.score}</td>
    </tr>
  {/each}
  </tbody>
</Table>

<style>
</style>
