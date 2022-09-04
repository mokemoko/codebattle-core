<script lang="ts">
  import { onMount } from 'svelte'
  import type {Match} from '../generated'

  export let match: Match
  export let showDetail = true

  let canvas: HTMLCanvasElement

  onMount(() => {
    const ctx = canvas.getContext('2d')
    ctx.fillStyle = "black"
    ctx.fillRect(0, 0, 720, 624)
  })

</script>

<canvas bind:this={canvas} width={720} height={624}></canvas>

<style>
</style>
