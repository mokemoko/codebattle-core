<script lang="ts">
  import { Badge } from 'sveltestrap'

  export let status = ''

  $: color = {
    // EntryStatus
    'accepted': 'primary',
    'processing': 'secondary',
    'registered': 'success',
    'error': 'danger',
    'disabled': 'dark',
    // MatchType
    'rated': 'primary',
    'unrated': 'secondary',
  }[status] || 'primary'
</script>

<Badge {color}>{status}</Badge>

<style>
</style>
