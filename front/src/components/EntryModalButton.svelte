<script lang="ts">
  import { Button, Form, FormGroup, Input, Modal, ModalBody, ModalFooter, ModalHeader } from 'sveltestrap'
  import { registerEntry } from '../domain/usecase'

  let isOpen = false
  let name = ''
  let repository = ''
  let error = ''

  const onSubmit = async () => {
    try {
      await registerEntry(name, repository)
    } catch (e) {
      error = e.message
      return
    }
    name = ''
    repository = ''
    isOpen = false
  }
</script>

<Button outline on:click={() => isOpen = true}>Entry</Button>

<Modal isOpen={isOpen}>
  <ModalHeader>Entry</ModalHeader>
  <ModalBody>
    <Form>
      <FormGroup floating label="Name">
        <Input bind:value={name}/>
      </FormGroup>
      <FormGroup floating label="Repository">
        <Input bind:value={repository}/>
      </FormGroup>
    </Form>
    {#if error.length > 0}
      <span class="text-danger">{error}</span>
    {/if}
  </ModalBody>
  <ModalFooter>
    <Button color="primary" on:click={onSubmit}>Submit</Button>
    <Button color="secondary" on:click={() => isOpen = false}>Cancel</Button>
  </ModalFooter>
</Modal>

<style>
</style>
