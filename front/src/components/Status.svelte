<script lang="ts">
  import { Badge, Button, Table } from 'sveltestrap'
  import type { User } from '../generated'

  export let user: User
</script>

<h3>
  <span>Your Status</span>
  <Button outline>Entry</Button>
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
  <tr>
    <td>1</td>
    <td>sample</td>
    <td>100</td>
    <td><Badge>Registered</Badge></td>
  </tr>
  </tbody>
</Table>

<style>
  h3 {
    display: flex;
    justify-content: space-between;
  }
</style>
