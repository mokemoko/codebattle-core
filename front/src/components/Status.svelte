<script lang="ts">
  import { Button, Table } from 'sveltestrap'
  import type { Entry } from '../generated'
  import EntryModal from './EntryModal.svelte'
  import { contestState, userState } from '../domain/state'
  import { get } from 'svelte/store'
  import StatusBadge from './atoms/StatusBadge.svelte'
  import ScoreLabel from './atoms/ScoreLabel.svelte'

  let entries: Entry[] = []
  let selectedEntry: Entry | null
  let isOpen = false

  contestState.subscribe(contest => {
    const userId = get(userState).id
    entries = contest.ranking.filter(entry => entry.user.id === userId)
  })

  const selectEntry = (entry: Entry | null) => {
    selectedEntry = entry
    isOpen = true
  }
</script>

<h3>
  <span>Your Status</span>
  <Button outline on:click={() => selectEntry(null)}>Entry</Button>
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
    <tr on:click={() => selectEntry(entry)}>
      <td>{i + 1}</td>
      <td>{entry.name}</td>
      <td>
        <ScoreLabel score={entry.score}/>
      </td>
      <td>
        <StatusBadge status={entry.status}/>
      </td>
    </tr>
  {/each}
  </tbody>
</Table>

<!-- model -->
<EntryModal isOpen={isOpen} entry={selectedEntry}
            name={selectedEntry?.name || ''} repository={selectedEntry?.repository || ''} isDisabled={selectedEntry?.status === "disabled"}
            callback={() => isOpen = false}/>

<style>
  h3 {
    display: flex;
    justify-content: space-between;
  }
</style>
