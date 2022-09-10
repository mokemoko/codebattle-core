<script lang="ts">
  import { onMount } from 'svelte'
  import type { Match } from '../generated'
  import { Game } from './game'
  import { Button, Icon, InputGroup } from 'sveltestrap'

  export let match: Match
  export let showDetail = true

  let canvas: HTMLCanvasElement
  let game: Game
  let turn: number
  let timer: NodeJS.Timer | null

  async function fetchLog(path: string) {
    const res = await fetch(path)
    return await res.text()
  }

  function switchAuto() {
    if (timer) {
      clearInterval(timer)
      timer = null
    } else {
      timer = setInterval(() => {
        turn += 1
      }, 100)
    }
  }

  function step(delta: number) {
    if (timer) {
      clearInterval(timer)
      timer = null
    }
    turn += delta
  }

  onMount(async () => {
    // TODO: fix
    // const log = await fetchLog(`logs/${match.id}.log`)
    const log = await fetchLog(`bomberman/log`)

    game = new Game(log, canvas)

    turn = 0
  })

  $: {
    if (typeof turn !== 'undefined') {
      game.draw(turn)
    }
  }

</script>

<div class="body d-inline-flex flex-column">
  <canvas bind:this={canvas} width={720} height={624}></canvas>
  <InputGroup>
    <Button outline color="primary" class="py-1 px-2" on:click={() => step(-1)} disabled={turn <= 0}>
      <Icon class="sm-icon" name="chevron-left"/>
    </Button>
    <Button outline color="primary" class="p-1" on:click={switchAuto}>
      {#if timer}
        <Icon class="md-icon" name="pause-fill"/>
      {:else}
        <Icon class="md-icon" name="play-fill"/>
      {/if}
    </Button>
    <Button outline color="primary" class="py-1 px-2 me-2" on:click={() => step(1)} disabled={turn >= game?.turnMax}>
      <Icon class="sm-icon" name="chevron-right"/>
    </Button>
    <input class="flex-grow-1" type="range" min="0" max={game?.turnMax || 0} step="1" bind:value={turn}/>
  </InputGroup>
</div>

<style>
  .body :global(.btn) {
    border: none
  }

  .body :global(.md-icon) {
    font-size: 32px
  }
  .body :global(.sm-icon) {
    font-size: 20px
  }
</style>
