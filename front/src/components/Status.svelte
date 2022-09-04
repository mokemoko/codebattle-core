<script lang="ts">
  import { Table } from 'sveltestrap'
  import type { User } from '../generated'

  export let user: User
</script>

<h3>Status</h3>
<Table striped hover>
  <thead>
  <tr>
    <th>#</th>
    <th>Name</th>
    <th>Score</th>
  </tr>
  </thead>
  <tbody>
  </tbody>
</Table>

<style>
</style>
