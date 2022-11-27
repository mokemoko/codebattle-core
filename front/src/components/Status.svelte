<script lang="ts">
  import { Badge, Table } from 'sveltestrap'
  import type { Entry } from '../generated'
  import EntryModalButton from './EntryModalButton.svelte'
  import { contestState, userState } from '../domain/state'
  import { get } from 'svelte/store'

  let entries: Entry[] = []

  contestState.subscribe(contest => {
    const userId = get(userState).id
    entries = contest.ranking.filter(entry => entry.user.id === userId)
  })
</script>

<h3>
  <span>Your Status</span>
  <EntryModalButton/>
</h3>
<Table striped hover>
  <thead>
  <tr>
    <th>#</th>
    <th>Name</th>
    <th>Score</th>
    <th>Status</th>
  </tr>
  </thead>
  <tbody>
  {#each entries as entry, i}
    <tr>
      <td>{i + 1}</td>
      <td>{entry.name}</td>
      <td>{entry.score}</td>
      <td>
        <Badge>{entry.status}</Badge>
      </td>
    </tr>
  {/each}
  </tbody>
</Table>

<style>
  h3 {
    display: flex;
    justify-content: space-between;
  }
</style>
