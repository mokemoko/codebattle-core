import { writable } from 'svelte/store'
// TODO: rollupで謎にエラーになるので暫定的にanyとする
// import type { User } from '../generated'

export const userState = writable()
