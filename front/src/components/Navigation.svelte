<script lang="ts">
  import {
    Collapse,
    Dropdown, DropdownItem, DropdownMenu, DropdownToggle,
    Nav, NavItem, NavLink, Navbar, NavbarBrand,
  } from 'sveltestrap'
  import { userState } from '../state/user'
  import type { User } from '../generated'

  let user: User | null

  userState.subscribe(u => user = u)

</script>

<Navbar color="light" light expand="md">
  <NavbarBrand href="#">Code Battle</NavbarBrand>
  <Collapse navbar isOpen>
    <Nav class="ms-auto" navbar>
      <Dropdown nav inNavbar>
        {#if user}
          <DropdownToggle nav caret>{user.name}</DropdownToggle>
          <DropdownMenu end>
            <DropdownItem href="#" disabled>Logout</DropdownItem>
          </DropdownMenu>
        {:else}
          <NavItem>
            <NavLink href="http://localhost:8081/login">Login</NavLink>
          </NavItem>
        {/if}
      </Dropdown>
    </Nav>
  </Collapse>
</Navbar>

<style>
</style>
