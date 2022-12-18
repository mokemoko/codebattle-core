<script lang="ts">
  import { Button, Form, FormGroup, Label, Modal, ModalBody, ModalFooter, ModalHeader } from 'sveltestrap'
  import { get } from 'svelte/store'
  import { contestState } from '../domain/state'
  import type { Entry } from '../generated'

  export let isOpen: boolean
  export let callback: () => {}

  let entries: Entry[] = get(contestState).ranking || []
  let entryIds: string[] = []
  let error = ''

  const validate = () => {
    if (entryIds.length !== 4) {
      throw Error('Choose exactly 4 entries.')
    }
  }

  const onSubmit = async () => {
    console.log(entryIds)
    try {
      validate()
      // TBA register match
    } catch (e) {
      error = e.message
      return
    }
    error = ''
    callback()
  }
</script>

<Modal isOpen={isOpen}>
  <ModalHeader>Create Match</ModalHeader>
  <ModalBody>
    <Form>
      <FormGroup>
        <Label for="entrySelect">Select target entries.</Label>
        <!-- sveltestrap が multiple select に対応していない -->
        <select class="form-select" id="entrySelect" multiple size={Math.min(entries.length, 10)} bind:value={entryIds}>
          {#each entries as entry}
            <option value={entry.id}>{entry.name} ({entry.user.name})</option>
          {/each}
        </select>
      </FormGroup>
    </Form>
    {#if error.length > 0}
      <span class="text-danger">{error}</span>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button color="primary" on:click={onSubmit}>Submit</Button>
    <Button color="secondary" on:click={callback}>Cancel</Button>
  </ModalFooter>
</Modal>

<style>
</style>
