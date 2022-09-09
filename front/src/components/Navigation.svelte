<script lang="ts">
  import {
    Collapse,
    Dropdown,
    DropdownItem,
    DropdownMenu, DropdownToggle,
    Nav,
    Navbar,
    NavbarBrand,
  } from 'sveltestrap'
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="#">Code Battle</NavbarBrand>
  <Collapse navbar isOpen>
    <Nav class="ms-auto" navbar>
      <Dropdown nav inNavbar>
        <DropdownToggle nav caret>Menu</DropdownToggle>
        <DropdownMenu end>
          <DropdownItem>Status</DropdownItem>
          <DropdownItem divider/>
          <DropdownItem>Logout</DropdownItem>
        </DropdownMenu>
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>

<style>
</style>
