<script lang="ts">
  import { Badge } from 'sveltestrap'

  export let status = ''

  $: color = {
    'processing': 'secondary',
    'registered': 'success',
    'error': 'danger',
    'disabled': 'dark',
  }[status] || 'primary'
</script>

<Badge {color}>{status}</Badge>

<style>
</style>
