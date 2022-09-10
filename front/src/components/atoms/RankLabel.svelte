<script lang="ts">

  export let rank: number | null

  $: label = rank ? (['1st', '2nd', '3rd'][rank - 1] || `${rank}th`) : '-'
</script>

<span class="font-monospace">{label}</span>

<style>
  span {
    white-space: pre;
  }
</style>
