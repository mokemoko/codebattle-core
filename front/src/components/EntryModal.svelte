<script lang="ts">
  import { Button, Form, FormGroup, FormText, Input, Modal, ModalBody, ModalFooter, ModalHeader } from 'sveltestrap'
  import { registerEntry, updateEntry } from '../domain/usecase'
  import type { Entry } from '../generated'

  export let isOpen: boolean
  export let entry: Entry | null
  export let callback: () => {}

  export let name = ''
  export let repository = ''
  export let isDisabled = false

  let error = ''

  const onSubmit = async () => {
    try {
      if (entry) {
        await updateEntry(entry.id, name, repository, isDisabled)
      } else {
        await registerEntry(name, repository, isDisabled)
      }
    } catch (e) {
      error = e.message
      return
    }
    name = ''
    repository = ''
    error = ''
    callback()
  }
</script>

<Modal isOpen={isOpen}>
  <ModalHeader>Entry</ModalHeader>
  <ModalBody>
    <Form>
      <FormGroup floating label="Name">
        <Input bind:value={name}/>
      </FormGroup>
      <FormGroup floating label="Repository">
        <Input bind:value={repository}/>
        <FormText>Dockerfileを直下に配置したリポジトリURLを入力してください。末尾にアンカーリンクでブランチ指定も可能です（#develop など）</FormText>
      </FormGroup>
      <FormGroup>
        <Input type="checkbox" label="非活性" bind:checked={isDisabled}/>
        <FormText>非活性状態のAIは対戦の対象とならず、ランキングにも表示されません</FormText>
      </FormGroup>
      {#if entry}
        <FormGroup floating label="Error">
          <Input value={entry.error} readonly/>
        </FormGroup>
      {/if}
    </Form>
    {#if error.length > 0}
      <span class="text-danger">{error}</span>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button color="primary" on:click={onSubmit}>{entry ? 'Update' : 'Register'}</Button>
    <Button color="secondary" on:click={callback}>Cancel</Button>
  </ModalFooter>
</Modal>

<style>
</style>
