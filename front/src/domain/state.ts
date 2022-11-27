import { writable } from 'svelte/store'
import type { User, Contest } from '../generated'

export const userState = writable<User>(null)

export const contestState = writable<Contest>(null)
