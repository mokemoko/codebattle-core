<script lang="ts">
  import { Button, Modal, ModalBody, ModalFooter, ModalHeader } from 'sveltestrap'

  export let isOpen: boolean
  export let callback: (isSubmitted: boolean) => {}

  const onSubmit = () => {
    callback(true)
  }
</script>

<Modal isOpen={isOpen}>
  <ModalHeader>Entry</ModalHeader>
  <ModalBody></ModalBody>
  <ModalFooter>
    <Button color="primary" on:click={onSubmit}>Submit</Button>
    <Button color="secondary" on:click={() => callback(false)}>Cancel</Button>
  </ModalFooter>
</Modal>

<style>
</style>
