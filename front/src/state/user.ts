import { writable } from 'svelte/store'
import type { User } from '../generated'

export const userState = writable<User>()
