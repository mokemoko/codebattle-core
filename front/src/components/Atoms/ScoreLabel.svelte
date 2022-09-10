<script lang="ts">

  export let score: number
  export let delta: number | null

  const colorMap: [number, string][] = [
    [1000, '#000'],
    [1500, '#500'],
    [2000, '#a00'],
    [Infinity, '#f00'],
  ]

  $: scoreColor = colorMap.find(m => score < m[0])[1]
</script>

<span class="font-monospace">
<span style="color: {scoreColor}">{score.toLocaleString()}</span>
  {#if delta}
  (<span class={delta >= 0 ? 'text-primary' : 'text-danger'}>{delta >= 0 ? '+' : ''}{delta.toLocaleString()}</span>)
{/if}
</span>

<style>
</style>
