import { Configuration, ContestsApi, UsersApi } from '../generated'

// TODO: ref env
const basePath = 'http://localhost:8081'

const baseConf = new Configuration({
  basePath,
  credentials: 'include',
})

export const userClient = new UsersApi(baseConf)
export const contestClient = new ContestsApi(baseConf)
export const entryClient = new ContestsApi(baseConf)
